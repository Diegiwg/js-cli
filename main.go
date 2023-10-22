package main

import (
	"os"

	"github.com/Diegiwg/js-cli/pkg/cmd"
	"github.com/Diegiwg/js-cli/pkg/console"
)

func main() {
	console.Clear()
	cmd.RegisterAllCommands()

	if len(os.Args) < 2 {
		cmd.Helper(&os.Args)
		os.Exit(1)
	}

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cmd.ExecuteCommand(os.Args[1], &args)
}
