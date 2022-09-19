package pgcommands_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	pg "github.com/habx/pg-commands"
	"github.com/habx/pg-commands/tests/fixtures"
)

func TestNewRestoreWrongCommand(t *testing.T) {
	savePGRestoreCmd := pg.PGRestoreCmd
	pg.PGRestoreCmd = "xxxx"
	Convey("Create new dump with wrong command", t, func() {
		restore, err := pg.NewRestore(fixtures.Setup())
		So(err, ShouldNotBeNil)
		So(restore, ShouldBeNil)
	})
	pg.PGRestoreCmd = savePGRestoreCmd
}
func TestNewRestore(t *testing.T) {
	restore, err := pg.NewRestore(fixtures.Setup())
	Convey("Create new restore", t, func() {
		So(err, ShouldBeNil)
		restore.SetPath("./")
		restore.SetSchemas([]string{"public"})
		So(restore.Options, ShouldNotBeEmpty)
		So(restore.Verbose, ShouldBeFalse)
		Convey("Public funcs", func() {
			restore.ResetOptions()
			So(restore.Options, ShouldBeEmpty)
			restore.EnableVerbose()
			So(restore.Verbose, ShouldBeTrue)
		})
	})
	Convey("Create without binary", t, func() {
		savePGRestoreCmd := pg.PGRestoreCmd
		pg.PGRestoreCmd = ""
		restore.ResetOptions()
		restoreBad := restore.Exec("test", pg.ExecOptions{StreamPrint: false})
		So(restoreBad.Error, ShouldNotBeNil)
		pg.PGRestoreCmd = savePGRestoreCmd
	})
}

func TestRestore(t *testing.T) {
	pgSetup := fixtures.Setup()
	dump, _ := pg.NewDump(pgSetup)
	result := dump.Exec(pg.ExecOptions{StreamPrint: false})
	Convey("Create standard restore", t, func() {
		restore, err := pg.NewRestore(fixtures.Setup())
		So(err, ShouldBeNil)
		x := restore.Exec(result.File, pg.ExecOptions{StreamPrint: true})
		So(x.Error, ShouldBeNil)
		So(x.FullCommand, ShouldNotBeEmpty)
		fmt.Println(x.FullCommand)
		So(x.FullCommand, ShouldEqual, fmt.Sprintf(
			"--no-owner --no-acl --clean --exit-on-error --dbname=%s --host=%s --port=%d --username=%s %s--schema=public %s",
			pgSetup.DB,
			pgSetup.Host,
			pgSetup.Port,
			pgSetup.Username,
			func() string {
				if restore.Role != "" {
					return fmt.Sprintf("--role=%s ", restore.Role)
				}

				return ""
			}(),
			result.File))
		restore.EnableVerbose()
		restore.Role = "dev_example"
		x = restore.Exec(result.File, pg.ExecOptions{StreamPrint: false})
		So(x.Error, ShouldBeNil)
		So(x.FullCommand, ShouldNotBeEmpty)
		So(x.FullCommand, ShouldEqual, fmt.Sprintf(
			"--no-owner --no-acl --clean --exit-on-error --dbname=%s --host=%s --port=%d --username=%s --role=%s -v --schema=public %s",
			pgSetup.DB,
			pgSetup.Host,
			pgSetup.Port,
			pgSetup.Username,
			restore.Role,
			result.File))
	})
	Convey("Create failed restore", t, func() {
		restore, err := pg.NewRestore(&pg.Postgres{})
		So(err, ShouldBeNil)
		result := restore.Exec("ok", pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldNotBeNil)
	})
}
