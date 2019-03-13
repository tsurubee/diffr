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
	flags.SetOutput(d.OutStream)

	flags.BoolVar(&version, "version", false, "Print version information")
	flags.BoolVar(&version, "v", false, "Print version information (short)")

	if err := flags.Parse(args[1:]); err != nil {
		helpInfo()
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(d.ErrStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	args = flags.Args()
	switch l := len(args); {
	case l == 0:
		fmt.Println("diffr: missing operand after `diffr`")
		helpInfo()
		return ExitCodeError
	case l == 1:
		fmt.Printf("diffr: missing operand after %v\n", args[0])
		helpInfo()
		return ExitCodeError
	case l == 2:
		return diffMain(args[0], args[1])
	case 2 < l:
		fmt.Printf("diffr: extra operand %v\n", args[2])
		helpInfo()
		return ExitCodeError
	}

	return ExitCodeOK
}
