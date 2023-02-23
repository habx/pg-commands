package pgcommands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func NewDefaultExecOptions() ExecOptions {
	return ExecOptions{}
}

func NewStreamToStdout() ExecOptions {
	return ExecOptions{StreamPrint: true, StreamDestination: os.Stdout}
}

func NewCustomExecOptions(streamPrint bool, dst io.Writer) ExecOptions {
	return ExecOptions{StreamPrint: streamPrint, StreamDestination: dst}
}

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
			return output, err
		}

		if options.StreamPrint {
			_, err = fmt.Fprintln(options.StreamDestination, line)
			if err != nil {
				return output, err
			}
		}

		output += line
	}
}

func CommandExist(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}
