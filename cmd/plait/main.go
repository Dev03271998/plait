// Command plait is a small declarative infrastructure-as-code engine.
package main

import (
	"os"

	"github.com/Dev03271998/plait/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], cli.IO{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}))
}
