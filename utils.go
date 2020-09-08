package pg_commands

import (
	"bufio"
	"fmt"
	"io"
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
			fmt.Printf(line)
		}
		line, err = reader.ReadString('\n')
		output += line
	}
	return output
}
