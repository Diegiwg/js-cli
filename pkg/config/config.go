package config

import (
	"os"
	"strings"
)

// Save the runtime and package manager in a config file insider the project root
func SaveConfig(rt string, pm string) {

	file, err := os.OpenFile(".js-cli-config", os.O_SYNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString("Runtime:" + rt + "\n" + "PackageManager:" + pm)
	if err != nil {
		panic(err)
	}

}

func LoadConfig() (string, string) {

	file, err := os.ReadFile(".js-cli-config")
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
