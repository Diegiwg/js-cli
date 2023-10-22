package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Diegiwg/js-cli/pkg/config"
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

type cmd struct {
	name  string
	desc  string
	usage string
}

var commands = []cmd{}

func Helper(args *[]string) {

	fmt.Println("Welcome to js-cli! Here's how you can use it:")
	fmt.Println("Usage: js-cli <command> [options]")

	fmt.Println("\nCommands:")
	for i := 0; i < len(commands); i++ {
		fmt.Printf("%s: %s\n", commands[i].name, commands[i].desc)
		fmt.Printf("    Usage: %s\n\n", commands[i].usage)
	}

	os.Exit(0)
}

func Init(args *[]string) {
	type initArgsType struct {
		Runtime        string
		PackageManager string
	}

	initArgs := initArgsType{}
	if *args != nil || len(*args) != 0 {
		// TODO: parse args
	}

	fmt.Println("Initializing js-cli...")

	if initArgs.Runtime == "" {
		var runtime string

		fmt.Println("Choose a runtime:")
		for i := 0; i < len(Runtime); i++ {
			fmt.Printf("%d: %s\n", i+1, Runtime[i])
		}

		fmt.Scanln(&runtime)

		switch runtime {
		case "1":
			initArgs.Runtime = Runtime[0]
		case "2":
			initArgs.Runtime = Runtime[1]
		case "3":
			initArgs.Runtime = Runtime[2]
		default:
			console.Clear()

			msg := "The runtime " + runtime + " is not supported."
			fmt.Println(msg)

			Helper(nil)
		}
	}

	if initArgs.PackageManager == "" {
		var packageManager string

		fmt.Println("Choose a package manager:")
		for i := 0; i < len(PackageManager); i++ {
			fmt.Printf("%d: %s\n", i+1, PackageManager[i])
		}

		fmt.Scanln(&packageManager)

		switch packageManager {
		case "1":
			initArgs.PackageManager = PackageManager[0]
		case "2":
			initArgs.PackageManager = PackageManager[1]
		case "3":
			initArgs.PackageManager = PackageManager[2]
		default:
			console.Clear()

			msg := "The package manager " + packageManager + " is not supported."
			fmt.Println(msg)

			Helper(nil)
		}
	}

	fmt.Println("js-cli initialized with the following values:")
	fmt.Println("Runtime: " + initArgs.Runtime)
	fmt.Println("PackageManager: " + initArgs.PackageManager)

	config.SaveConfig(initArgs.Runtime, initArgs.PackageManager)

	state.Runtime = initArgs.Runtime
	state.PackageManager = initArgs.PackageManager

	cmd := state.PackageManager + " init -y"
	console.Exec(cmd)

}

func Install(args *[]string) {
	console.Clear()

	cmd := state.PackageManager + " install " + strings.Join(*args, " ")
	console.Exec(cmd)
}

func List() {
	console.Clear()

	console.Exec(state.PackageManager + " list")
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
