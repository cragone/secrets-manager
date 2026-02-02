package main

import (
	"flag"
	"fmt"
	"manager/actions"
	"os"
)

func main() {
	actions.Init()
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		_ = actions.Help(args)
		os.Exit(1)
	}

	cmdKey := actions.Command(args[0])
	cmd, ok := actions.Commands[cmdKey]
	if !ok {
		_ = actions.Help(args)
		os.Exit(1)

	}
	if err := cmd(args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
