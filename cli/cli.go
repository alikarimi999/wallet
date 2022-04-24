package cli

import (
	"fmt"
	"os"
	"runtime"
)

type Commandline struct{}

func (cli *Commandline) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		runtime.Goexit()
	}
}

func (cli *Commandline) PrintUsage() {
	fmt.Println("Usage:")
}
