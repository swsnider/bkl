package bkl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	path string
	docs []any
}

var baseTemplate = ""

func loadFile(path string) (*file, error) {
	f := &file{
		path: path,
	}

	format, found := formatByExtension[ext(path)]
	if !found {
		return nil, fmt.Errorf("%s: %w", path, ErrUnknownFormat)
	}

	fh, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	defer fh.Close()

	raw, err := io.ReadAll(fh)
	if err != nil {
		return nil, err
	}

	rawDocs := bytes.SplitAfter(raw, []byte("\n---\n"))

	for i, rawDoc := range rawDocs {
		// Leave the initial \n attached
		rawDoc = bytes.TrimSuffix(rawDoc, []byte("---\n"))

		doc, err := format.decode(rawDoc)
		if err != nil {
			return nil, ErrorsJoin(fmt.Errorf("%s[doc%d]: %w", path, i, ErrDecode), err)
		}

		doc = normalize(doc)
		doc = env(doc)

		f.docs = append(f.docs, doc)
	}

	return f, nil
}

func (f *file) parent() (*string, error) {
	parent, err := f.parentFromDirective()
	if err != nil {
		return nil, err
	}

	if parent != nil {
		return parent, nil
	}

	parent, err = f.parentFromSymlink()
	if err != nil {
		return nil, err
	}

	if parent != nil {
		return parent, nil
	}

	return f.parentFromFilename()
}

func (f *file) parentFromDirective() (*string, error) {
	docMap, ok := f.docs[0].(map[string]any)
	if !ok {
		return nil, nil
	}

	parent, found := docMap["$parent"]
	if !found {
		return nil, nil
	}

	delete(docMap, "$parent")

	if parent == nil {
		return &baseTemplate, nil
	}

	parentStr, ok := parent.(string)
	if !ok {
		return nil, fmt.Errorf("%T: %w", parent, ErrInvalidParentType)
	}

	parentStr = filepath.Join(filepath.Dir(f.path), parentStr)

	parentPath := findFile(parentStr)
	if parentPath == "" {
		return nil, fmt.Errorf("%s: %w", parentStr, ErrMissingFile)
	}

	return &parentPath, nil
}

func (f *file) parentFromSymlink() (*string, error) {
	dest, err := filepath.EvalSymlinks(f.path)
	if err != nil {
		return nil, err
	}

	if dest == f.path {
		// Not a link
		return nil, nil
	}

	f.path = dest

	return f.parentFromFilename()
}

func (f *file) parentFromFilename() (*string, error) {
	dir := filepath.Dir(f.path)
	base := filepath.Base(f.path)

	parts := strings.Split(base, ".")
	// Last part is file extension

	switch {
	case len(parts) < 2:
		return nil, fmt.Errorf("[%s] %w", f.path, ErrInvalidFilename)

	case len(parts) == 2:
		return &baseTemplate, nil

	default:
		layerPath := filepath.Join(dir, strings.Join(parts[:len(parts)-2], "."))

		extPath := findFile(layerPath)
		if extPath == "" {
			return nil, fmt.Errorf("[%s]: %w", layerPath, ErrMissingFile)
		}

		return &extPath, nil
	}
}
