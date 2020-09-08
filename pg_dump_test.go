package pg_commands_test

import (
	"testing"

	"github.com/habx/pg-commands/tests/fixtures"

	pg "github.com/habx/pg-commands"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDump(t *testing.T) {
	dump := pg.NewDump(fixtures.Setup())
	Convey("Create new dump", t, func() {
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
		dump := pg.NewDump(fixtures.Setup())
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
	})
	Convey("Create dump with ignore table", t, func() {
		dump := pg.NewDump(fixtures.Setup())
		So(dump.IgnoreTableDataToString(), ShouldBeEmpty)
		dump.IgnoreTableData = append(dump.IgnoreTableData, "public.test_1")
		So(dump.IgnoreTableDataToString(), ShouldNotBeEmpty)
		dump.IgnoreTableData = append(dump.IgnoreTableData, "public.test_1")
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
	})
	Convey("Create dump with log and custom format", t, func() {
		dump := pg.NewDump(fixtures.Setup())
		dump.EnableVerbose()
		dump.SetupFormat("t")
		result := dump.Exec(pg.ExecOptions{StreamPrint: true})
		So(result.Error, ShouldBeNil)
		So(result.FullCommand, ShouldNotBeEmpty)
		So(result.File, ShouldNotBeEmpty)
		So(result.Mine, ShouldEqual, "application/x-tar")
	})
	Convey("Create failed dump", t, func() {
		dump := pg.NewDump(&pg.Postgres{})
		result := dump.Exec(pg.ExecOptions{StreamPrint: false})
		So(result.Error, ShouldNotBeNil)
	})
}
