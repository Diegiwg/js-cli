package main

import (
	"os"

	"github.com/Diegiwg/js-cli/pkg/cmd"
	"github.com/Diegiwg/js-cli/pkg/config"
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

func command(arg string, args *[]string) {

	state.CurrentDir, _ = os.Getwd()
	state.CurrentConfigFile = state.CurrentDir + "\\.js-cli-config"
	state.Runtime, state.PackageManager = config.LoadConfig()

	if state.Runtime == "" {
		state.Runtime = cmd.Runtime[0]
	}

	if state.PackageManager == "" {
		state.PackageManager = cmd.PackageManager[0]
	}

	switch arg {
	case "init", "i":
		{
			cmd.Init(args)
		}

	case "install", "add":
		{
			cmd.Install(args)
		}

	case "list", "l":
		{
			cmd.List()
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
