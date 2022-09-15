package pg_commands_test

import (
	"fmt"
	"testing"

	"github.com/habx/pg-commands/tests/fixtures"

	pg "github.com/habx/pg-commands"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDumpWrongCommand(t *testing.T) {
	savePGDumpCmd := pg.PGDumpCmd
	pg.PGDumpCmd = "xxxx"
	Convey("Create new dump with wrong command", t, func() {
		dump, err := pg.NewDump(fixtures.Setup())
		So(err, ShouldNotBeNil)
		So(dump, ShouldBeNil)
	})
	pg.PGDumpCmd = savePGDumpCmd
}

func TestNewDump(t *testing.T) {
	dump, err := pg.NewDump(fixtures.Setup())
	Convey("Create new dump", t, func() {
		So(err, ShouldBeNil)
		So(dump.Options, ShouldNotBeEmpty)
		So(dump.Verbose, ShouldBeFalse)
		Convey("Public funcs", func() {
			dump.ResetOptions()
			So(dump.Options, ShouldBeEmpty)
			dump.EnableVerbose()
			So(dump.Verbose, ShouldBeTrue)
		})
	})
}

func TestDump(t *testing.T) {
	Convey("Create standard dump", t, func() {
		pgSetup := fixtures.Setup()
		dump, err := pg.NewDump(pgSetup)
		So(err, ShouldBeNil)
		dump.SetFileName("test-dump.sql.tar.gz")
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
		So(result.FullCommand, ShouldEqual, fmt.Sprintf(
			"--no-owner --no-acl --clean --blob --dbname=%s --host=%s --port=%d --username=%s -Fc -f%s",
			pgSetup.DB,
			pgSetup.Host,
			pgSetup.Port,
			pgSetup.Username,
			result.File))
	})
	Convey("Create dump with ignore table", t, func() {
		pgSetup := fixtures.Setup()
		dump, err := pg.NewDump(pgSetup)
		So(err, ShouldBeNil)
		So(dump.IgnoreTableDataToString(), ShouldBeEmpty)
		dump.IgnoreTableData = append(dump.IgnoreTableData, "public.test_1")
		So(dump.IgnoreTableDataToString(), ShouldNotBeEmpty)
		dump.IgnoreTableData = append(dump.IgnoreTableData, "public.test_1")
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
		So(result.FullCommand, ShouldEqual, fmt.Sprintf(
			"--no-owner --no-acl --clean --blob --dbname=%s --host=%s --port=%d --username=%s -Fc --exclude-table-data=public.test_1 --exclude-table-data=public.test_1 -f%s",
			pgSetup.DB,
			pgSetup.Host,
			pgSetup.Port,
			pgSetup.Username,
			result.File))
	})
	Convey("Create dump with log and custom format", t, func() {
		pgSetup := fixtures.Setup()
		dump, err := pg.NewDump(pgSetup)
		So(err, ShouldBeNil)
		dump.EnableVerbose()
		dump.SetupFormat("t")
		dump.SetPath("./")
		result := dump.Exec(pg.ExecOptions{StreamPrint: true})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
		So(result.FullCommand, ShouldEqual, fmt.Sprintf(
			"--no-owner --no-acl --clean --blob --dbname=%s --host=%s --port=%d --username=%s -Ft -v -f%s",
			pgSetup.DB,
			pgSetup.Host,
			pgSetup.Port,
			pgSetup.Username,
			dump.Path+result.File))
	})
	Convey("Create failed dump", t, func() {
		dump, err := pg.NewDump(&pg.Postgres{})
		So(err, ShouldBeNil)
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldNotBeNil)
		So(result.FullCommand, ShouldEqual, fmt.Sprintf("--no-owner --no-acl --clean --blob -Fc -f%s", result.File))
	})
}
