package pgcommands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

type ExecOptions struct {
	StreamPrint       bool
	StreamDestination io.Writer
}

func streamExecOutput(out io.ReadCloser, options ExecOptions) (string, error) {
	output := ""
	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return output, nil
			}

			return output, fmt.Errorf("error reading output: %w", err)
		}

		if options.StreamPrint {
			_, err = fmt.Fprint(options.StreamDestination, line)
			if err != nil {
				return output, fmt.Errorf("error writing output: %w", err)
			}
		}

		output += line
	}
}

func streamOutput(stderrIn io.ReadCloser, opts ExecOptions, result *Result) {
	output, err := streamExecOutput(stderrIn, opts)
	if err != nil {
		result.Error = &ResultError{Err: err, CmdOutput: output}
	}
	result.Output = output
}

func CommandExist(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}
