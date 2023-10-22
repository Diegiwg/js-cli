package cmd

import (
	"os"

	"github.com/Diegiwg/js-cli/pkg/config"
	"github.com/Diegiwg/js-cli/pkg/state"
)

var Runtime = []string{"node", "deno", "bun"}
var PackageManager = []string{"npm", "pnpm", "yarn"}

type cmd struct {
	name  string
	desc  string
	usage string
}

var commands = []cmd{}

func initialize() {
	state.CurrentDir, _ = os.Getwd()
	state.CurrentConfigFile = state.CurrentDir + "\\.js-cli-config"
	state.Runtime, state.PackageManager = config.LoadConfig()

	if state.Runtime == "" {
		state.Runtime = Runtime[0]
	}

	if state.PackageManager == "" {
		state.PackageManager = PackageManager[0]
	}
}

func RegisterAllCommands() {
	commands = append(commands, cmd{
		name:  "init|i",
		desc:  "Initialize js-cli in your project.",
		usage: "js-cli init [options]\n\t--rt <runtime>\n\t--pm <package-manager>",
	})

	commands = append(commands, cmd{
		name:  "install|add",
		desc:  "Install a package in your project.",
		usage: "js-cli install <package>",
	})

	commands = append(commands, cmd{
		name:  "list|l",
		desc:  "List all installed packages.",
		usage: "js-cli list",
	})
}

func ExecuteCommand(command string, args *[]string) {
	initialize()

	switch command {
	case "init", "i":
		{
			Init(args)
		}

	case "install", "add":
		{
			Install(args)
		}

	case "list", "l":
		{
			List()
		}

	default:
		{
			Helper(args)
		}
	}
}
