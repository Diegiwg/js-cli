package cmd

import (
	"flag"
	"fmt"

	"github.com/Diegiwg/js-cli/pkg/config"
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
var runtimeArg = initCmd.String("rt", "", "Runtime")
var packageManagerArg = initCmd.String("pm", "", "Package Manager")

func Init(args *[]string) {
	initCmd.Parse(*args)

	type initArgsType struct {
		Runtime        string
		PackageManager string
	}

	initArgs := initArgsType{}
	if *runtimeArg != "" {
		initArgs.Runtime = *runtimeArg
	}

	if *packageManagerArg != "" {
		initArgs.PackageManager = *packageManagerArg
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
