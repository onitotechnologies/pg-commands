package pgcommands_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	pg "github.com/onitotechnologies/pg-commands"
)

func CommandExistTest(t *testing.T) {
	Convey("Comand exist", t, func() {
		So(pg.CommandExist("xxxx"), ShouldBeFalse)
	})
}
