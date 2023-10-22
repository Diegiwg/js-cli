package cmd

import (
	"strings"

	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

// ! In this case, the command note receive any flags.

func Install(args *[]string) {
	console.Clear()

	cmd := state.PackageManager + " install " + strings.Join(*args, " ")
	console.Exec(cmd)
}
