package main

import (
	"os"

	"github.com/Diegiwg/js-cli/pkg/cmd"
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

func command(arg string, args *[]string) {

	state.Init()

	switch arg {
	case "init", "i":
		{
			cmd.Init(args)
		}

	case "install", "add":
		{
			cmd.Install(args)
		}

	default:
		{
			cmd.Helper(args)
		}
	}

}

func main() {
	console.Clear()
	cmd.RegisterAllCommands()

	if len(os.Args) < 2 {
		cmd.Helper(&os.Args)
		os.Exit(1)
	}

	cmd := os.Args[1]

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	command(cmd, &args)
}
