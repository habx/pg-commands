package pg_commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	// PGRestoreCmd is the path to the `pg_restore` executable
	PGRestoreCmd      = "pg_restore"
	PGDRestoreStdOpts = []string{"--no-owner", "--no-acl", "--clean", "--schema=public", "--exit-on-error"}
)

type Restore struct {
	*Postgres
	// Verbose mode
	Verbose bool
	// Role: do SET ROLE before restore
	Role string
	// Path: setup path for source restore
	Path string
	// Extra pg_dump options
	// e.g []string{"--inserts"}
	Options []string
}

func NewRestore(pg *Postgres) *Restore {
	return &Restore{Options: PGDRestoreStdOpts, Postgres: pg}
}

// Exec `pg_restore` of the specified database, and restore from a gzip compressed tarball archive.
func (x *Restore) Exec(filename string) Result {
	result := Result{}
	options := append(x.restoreOptions(), fmt.Sprintf("%s%s", x.Path, filename))
	result.FullCommand = strings.Join(options, " ")
	cmd := exec.Command(PGRestoreCmd, options...)

	cmd.Env = append(os.Environ(), x.EnvPassword)
	out, err := cmd.CombinedOutput()
	if err != nil {
		result.Error = &ResultError{Err: err, CmdOutput: string(out)}
		if exitError, ok := err.(*exec.ExitError); ok {
			result.Error.ExitCode = exitError.ExitCode()
		}
	}
	result.Output = string(out)
	return result
}

func (x *Restore) ResetOptions() {
	x.Options = []string{}
}

func (x *Restore) EnableVerbose() {
	x.Verbose = true
}

func (x *Restore) SetPath(path string) {
	x.Path = path
}

func (x *Restore) restoreOptions() []string {
	options := x.Options
	options = append(options, x.Postgres.Parse()...)

	if x.Role != "" {
		options = append(options, fmt.Sprintf(`--role=%v`, x.Role))
	} else if x.DB != "" {
		options = append(options, fmt.Sprintf(`--role=%v`, x.DB))
	}

	if x.Verbose {
		options = append(options, "-v")
	}
	return options
}
