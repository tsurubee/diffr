package diffr

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

type Diffr struct {
	OutStream io.Writer
	ErrStream io.Writer
}

func helpInfo() {
	fmt.Println("diffr: Try `diffr -help' for more information.")
}

func (d *Diffr) Run(args []string) int {
	var (
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(d.ErrStream)

	flags.BoolVar(&version, "v", false, "Print version information.")
	flags.BoolVar(&version, "version", false, "Print version information.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		helpInfo()
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(d.ErrStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}
