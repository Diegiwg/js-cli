package cmd

var Runtime = []string{"node", "deno", "bun"}
var PackageManager = []string{"npm", "pnpm", "yarn"}

type cmd struct {
	name  string
	desc  string
	usage string
}

var commands = []cmd{}

func RegisterAllCommands() {
	commands = append(commands, cmd{
		name:  "init|i",
		desc:  "Initialize js-cli in your project.",
		usage: "js-cli init [options]\n\t--rt <runtime>\n\t--pm <package-manager>",
	})

	commands = append(commands, cmd{
		name:  "install|add",
		desc:  "Install a package in your project.",
		usage: "js-cli install <package>",
	})

	commands = append(commands, cmd{
		name:  "list|l",
		desc:  "List all installed packages.",
		usage: "js-cli list",
	})
}
