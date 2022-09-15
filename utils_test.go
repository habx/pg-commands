package pg_commands_test

import (
	"testing"

	pg "github.com/habx/pg-commands"

	. "github.com/smartystreets/goconvey/convey"
)

func CommandExistTest(t *testing.T) {
	Convey("Comand exist", t, func() {
		So(pg.CommandExist("xxxx"), ShouldBeFalse)
	})
}
