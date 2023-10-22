package cmd

import (
	"github.com/Diegiwg/js-cli/pkg/console"
	"github.com/Diegiwg/js-cli/pkg/state"
)

func List() {
	console.Clear()

	console.Exec(state.PackageManager + " list")
}
