package state

import "github.com/Diegiwg/js-cli/pkg/config"

var Runtime string
var PackageManager string

func Init() {

	Runtime, PackageManager = config.LoadConfig()
}
