package pg_commands_test

import (
	"testing"

	"github.com/habx/pg-commands/tests/fixtures"

	pg "github.com/habx/pg-commands"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRestore(t *testing.T) {
	restore := pg.NewRestore(fixtures.Setup())
	Convey("Create new restore", t, func() {
		So(restore.Options, ShouldNotBeEmpty)
		So(restore.Verbose, ShouldBeFalse)
		Convey("Public funcs", func() {
			restore.ResetOptions()
			So(restore.Options, ShouldBeEmpty)
			restore.EnableVerbose()
			So(restore.Verbose, ShouldBeTrue)
		})
	})
}

func TestRestore(t *testing.T) {
	dump := pg.NewDump(fixtures.Setup())
	result := dump.Exec()
	Convey("Create standard restore", t, func() {
		restore := pg.NewRestore(fixtures.Setup())
		x := restore.Exec(result.File)
		So(x.Error, ShouldBeNil)
		So(x.FullCommand, ShouldNotBeEmpty)

		restore.EnableVerbose()
		restore.Role = "dev_example"
		x = restore.Exec(result.File)
		So(x.Error, ShouldBeNil)
		So(x.FullCommand, ShouldNotBeEmpty)
	})
	Convey("Create failed restore", t, func() {
		restore := pg.NewRestore(&pg.Postgres{})
		result := restore.Exec("ok")
		So(result.Error, ShouldNotBeNil)
	})
}
