package cmd

import (
	"strings"

	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

func Install(args *[]string) {
	console.Clear()

	cmd := state.PackageManager + " install " + strings.Join(*args, " ")
	console.Exec(cmd)
}
