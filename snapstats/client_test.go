package snapstats

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("When creating client", t, func() {
		c, err := NewClient("http://localhost:8181", true)

		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
		So(c.http, ShouldNotBeNil)
		So(c.prefix, ShouldEqual, "http://localhost:8181")
	})
}
