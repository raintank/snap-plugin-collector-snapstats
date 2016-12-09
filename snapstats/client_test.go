package snapstats

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	SkipConvey("When creating client", t, func() {
		c, err := NewClient("http://localhost:8181", true)

		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
		So(c.http, ShouldNotBeNil)
		So(c.prefix, ShouldEqual, "http://localhost:8181/v1")

		tasks, err := c.GetTasks()
		So(err, ShouldBeNil)
		So(tasks[0].TaskState, ShouldEqual, "Disabled")
		So(tasks[0].HitCount, ShouldEqual, 10)
		So(tasks[0].FailedCount, ShouldEqual, 10)
	})
}
