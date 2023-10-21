package cmd

import "fmt"

func Init(args *[]string) {
	type initArgsType struct {
		Runtime        string
		PackageManager string
	}

	initArgs := initArgsType{}
	if *args != nil || len(*args) != 0 {
		// TODO: parse args
	}

	if initArgs.Runtime == "" {
		initArgs.Runtime = "node"
		// TODO: Ask user for runtime
	}

	if initArgs.PackageManager == "" {
		initArgs.PackageManager = "npm"
		// TODO: Ask user for package manager
	}

	fmt.Println("Runtime: " + initArgs.Runtime)
	fmt.Println("PackageManager: " + initArgs.PackageManager)

	// config.SaveConfig(initArgs.Runtime, initArgs.PackageManager)
}
