package main

import (
	"github.com/Anas02200/goChain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
