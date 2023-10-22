package cmd

import (
	"fmt"
	"os"
)

// ! In this case, the command note receive any flags.

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
