package config

import (
	"os"
	"strings"

	"github.com/Diegiwg/js-cli/pkg/state"
)

// Save the runtime and package manager in a config file insider the project root
func SaveConfig(rt string, pm string) {

	// try to remove the old config file
	_ = os.Remove(state.CurrentConfigFile)

	file, err := os.OpenFile(state.CurrentConfigFile, os.O_SYNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	options := "Runtime:" + rt + "\n" + "PackageManager:" + pm
	_, err = file.WriteString(options)
	if err != nil {
		panic(err)
	}

}

func LoadConfig() (string, string) {
	_, err := os.Stat(state.CurrentConfigFile)
	if os.IsNotExist(err) {
		return "", ""
	}

	file, err := os.ReadFile(state.CurrentConfigFile)
	if err != nil {
		panic(err)
	}

	strFile := string(file)

	// Search for Runtime and PackageManager
	if !strings.Contains(strFile, "Runtime") || !strings.Contains(strFile, "PackageManager") {
		panic("Runtime or Package Manager not found in .js-cli-config")
	}

	rt := strings.Split(strings.Split(strFile, "Runtime:")[1], "\n")[0]
	pm := strings.Split(strings.Split(strFile, "PackageManager:")[1], "\n")[0]

	if rt == "" || pm == "" {
		panic("Runtime or Package Manager not found in .js-cli-config")
	}

	return rt, pm
}
