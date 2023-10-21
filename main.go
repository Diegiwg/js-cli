package main

import (
	"fmt"
	"os"

	"github.com/Diegiwg/js-cli/pkg/cmd"
	"github.com/Diegiwg/js-cli/pkg/config"
)

func command(arg string, args *[]string) bool {

	switch arg {
	case "init", "i":
		{
			cmd.Init(args)
		}

	default:
		{
			fmt.Println("Unknown command: " + arg)
		}
	}

	return true
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: js-cli <command> [args]")
		os.Exit(1)
	}

	cmd := os.Args[1]

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	command(cmd, &args)

	x, y := config.LoadConfig()
	fmt.Println(x, y)
}
