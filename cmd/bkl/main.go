package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gopatchy/bkl"
	"github.com/jessevdk/go-flags"
)

type options struct {
	OutputPath   *string `short:"o" long:"output" description:"output file path"`
	OutputFormat *string `short:"f" long:"format" description:"output format"`
	SkipParent   bool    `short:"P" long:"skip-parent" description:"skip loading parent templates"`
	Verbose      bool    `short:"v" long:"verbose" description:"enable verbose logging"`

	Positional struct {
		InputPaths []string `positional-arg-name:"inputPath" required:"1" description:"input file path"`
	} `positional-args:"yes"`
}

func main() {
	if os.Getenv("BKL_VERSION") != "" {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fatal(fmt.Errorf("ReadBuildInfo() failed")) //nolint:goerr113
		}

		fmt.Printf("%s", bi)
		os.Exit(0)
	}

	opts := &options{}

	fp := flags.NewParser(opts, flags.Default)

	_, err := fp.Parse()
	if err != nil {
		os.Exit(1)
	}

	p := bkl.New()

	if opts.Verbose {
		p.SetDebug(true)
	}

	format := ""
	if opts.OutputFormat != nil {
		format = *opts.OutputFormat
	}

	for _, path := range opts.Positional.InputPaths {
		fileP := bkl.New()

		if opts.Verbose {
			fileP.SetDebug(true)
		}

		realPath, f, err := bkl.FileMatch(path)
		if err != nil {
			fatal(err)
		}

		if format == "" && opts.OutputPath == nil {
			format = f
		}

		if opts.SkipParent {
			err = fileP.MergeFile(realPath)
		} else {
			err = fileP.MergeFileLayers(realPath)
		}

		if err != nil {
			fatal(err)
		}

		err = p.MergeParser(fileP)
		if err != nil {
			fatal(err)
		}
	}

	if opts.OutputPath == nil {
		err = p.OutputToWriter(os.Stdout, format)
	} else {
		err = p.OutputToFile(*opts.OutputPath, format)
	}

	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
