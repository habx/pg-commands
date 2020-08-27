package main

import (
	"fmt"
	"os/exec"
)

var commands = []string{
	"apt-get install -y lsb-release",
	"echo \"deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main\" > /etc/apt/sources.list.d/pgdg.list",
	"wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -",
	"apt-get -y update",
	"apt-get install -y postgresql-client",
}

func main() {
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
