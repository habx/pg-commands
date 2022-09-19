package deps

import (
	"fmt"
	"os/exec"
)

var commands = []string{
	"sudo apt-get install -y lsb-release",
	"sudo echo \"deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main\" | sudo tee -a /etc/apt/sources.list.d/pgdg.list",
	"wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -",
	"sudo apt-get -y update",
	"sudo apt-get install -y postgresql-client",
}

func InstallCommands() {
	for _, command := range commands {
		Command(command)
	}
}

func Command(c string) {
	out, err := exec.Command("/bin/bash", "-c", c).CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Print(exitError.ExitCode())
		}
		fmt.Print(string(out))
		panic(err)
	}
	fmt.Print(string(out))
}
