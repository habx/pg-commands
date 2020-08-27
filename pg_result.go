package pg_commands

type Results struct {
	Dump    Result
	Restore Result
}

type Result struct {
	Mine        string
	File        string
	Output      string
	Error       *ResultError
	FullCommand string
}

type ResultError struct {
	Err       error
	CmdOutput string
	ExitCode  int
}
