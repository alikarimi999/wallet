package main

import (
	"os"

	"github.com/alikarimi999/wallet/cli"
)

func main() {
	defer os.Exit(0)
	cmd := new(cli.Commandline)
	cmd.Run()
}
