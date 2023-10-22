package console

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "darwin":
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Exec(action string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", action)
	case "darwin":
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
