// Package bkl implements a layered configuration language parser.
package bkl

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Err              = fmt.Errorf("bkl error")
	ErrUnknownFormat = fmt.Errorf("unknown format (%w)", Err)
	ErrDecode        = fmt.Errorf("decoding error (%w)", Err)
)

// Parser carries state for parse operations with multiple layered inputs.
type Parser struct {
	docs  []any
	debug bool
}

// New creates and returns a new [Parser] with an empty starting document.
//
// New always succeeds and returns a Parser instance.
func New() *Parser {
	return &Parser{}
}

// NewFromFile creates a new [Parser] then calls [MergeFileLayers()] with
// the supplied path.
func NewFromFile(path string) (*Parser, error) {
	p := New()

	err := p.MergeFileLayers(path)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Parser) SetDebug(debug bool) {
	p.debug = debug
}

// MergePatch applies the supplied patch to the [Parser]'s current internal
// document state (at the specified document index) using bkl's merge
// semantics.
func (p *Parser) MergePatch(index int, patch any) error {
	if index >= len(p.docs) {
		p.docs = append(p.docs, make([]any, index-len(p.docs)+1)...)
	}

	// XXX
	p.docs[index] = patch

	return nil
}

// MergeIndexBytes parses the supplied doc bytes as the format specified by ext
// (file extension), then calls [MergePatch()].
func (p *Parser) MergeIndexBytes(index int, doc []byte, ext string) error {
	f, found := formatByExtension[ext]
	if !found {
		return fmt.Errorf("%s: %w", ext, ErrUnknownFormat)
	}

	patch, err := f.decode(doc)
	if err != nil {
		return fmt.Errorf("%w / %w", err, ErrDecode)
	}

	err = p.MergePatch(index, patch)
	if err != nil {
		return err
	}

	return nil
}

// MergeMultiBytes calls [MergeIndexBytes()] once for each item in the outer
// slice.
func (p *Parser) MergeMultiBytes(bs [][]byte, ext string) error {
	for i, b := range bs {
		err := p.MergeIndexBytes(i, b, ext)
		if err != nil {
			return fmt.Errorf("document index %d (of [0,%d]): %w", i, len(bs)-1, err)
		}
	}

	return nil
}

// MergeBytes splits its input into multiple documents (using the ---
// delimiter) then calls [MergeMultiBytes()].
func (p *Parser) MergeBytes(b []byte, ext string) error {
	docs := bytes.SplitAfter(b, []byte("\n---\n"))

	for i, doc := range docs {
		// Leave the initial \n attached
		docs[i] = bytes.TrimSuffix(doc, []byte("---\n"))
	}

	return p.MergeMultiBytes(docs, ext)
}

// MergeReader reads all input then calls [MergeBytes()].
func (p *Parser) MergeReader(in io.Reader, ext string) error {
	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	return p.MergeBytes(b, ext)
}

// MergeFile opens the supplied path and determines the file format from the
// file extension, then calls [MergeReader()].
func (p *Parser) MergeFile(path string) error {
	p.log("loading %s", path)

	fh, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	defer fh.Close()

	parts := strings.Split(filepath.Base(path), ".")
	ext := parts[len(parts)-1]

	err = p.MergeReader(fh, ext)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	return nil
}

// MergeFileLayers determines relevant layers from the supplied path and merges
// them in order.
func (p *Parser) MergeFileLayers(path string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	parts := strings.Split(base, ".")
	ext := parts[len(parts)-1]

	for i := 1; i < len(parts); i++ {
		layerParts := []string{}
		layerParts = append(layerParts, parts[:i]...)
		layerParts = append(layerParts, ext)

		layerPath := filepath.Join(dir, strings.Join(layerParts, "."))

		dest, _ := os.Readlink(layerPath)
		if dest != "" {
			err := p.MergeFileLayers(dest)
			if err != nil {
				return err
			}

			continue
		}

		err := p.MergeFile(layerPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) log(format string, v ...any) {
	if !p.debug {
		return
	}

	log.Printf(format, v...)
}
