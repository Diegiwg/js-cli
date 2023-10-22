package cmd

import (
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

// ! In this case, the command note receive any flags.

func List() {
	console.Clear()

	console.Exec(state.PackageManager + " list")
}
