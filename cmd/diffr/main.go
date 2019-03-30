package main

import (
	"os"
	"github.com/tsurubee/diffr"
)

func main() {
	d := diffr.Diffr{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(d.Run(os.Args))
}
