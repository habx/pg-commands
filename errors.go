package pgcommands

import "fmt"

type ErrCommandNotFound struct {
	Command string
}

func (e ErrCommandNotFound) Error() string {
	return fmt.Sprintf("command not found: %s", e.Command)
}
