package pgcommands

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type ExecOptions struct {
	StreamPrint bool
}

func streamExecOutput(out io.ReadCloser, options ExecOptions) string {
	output := ""
	reader := bufio.NewReader(out)
	line, err := reader.ReadString('\n')
	output += line
	for err == nil {
		if options.StreamPrint {
			//nolint: staticcheck
			fmt.Printf(line)
		}
		line, err = reader.ReadString('\n')
		output += line
	}

	return output
}
func CommandExist(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}
