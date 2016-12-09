package snapstats

import (
	"fmt"
	"strings"
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

var cases = []struct {
	tasks    []Task
	metrics  []plugin.Metric
	expected []string
}{
	{
		tasks:   []Task{},
		metrics: getTotalStatsMetricTypes(),
		expected: []string{
			"grafanalabs.snapstats.tasks.state.Disabled.count 0",
			"grafanalabs.snapstats.tasks.state.Running.count 0",
			"grafanalabs.snapstats.tasks.state.Stopped.count 0",
			"grafanalabs.snapstats.tasks.hitcount 0",
			"grafanalabs.snapstats.tasks.failedcount 0",
		},
	},
	{
		tasks: []Task{
			{
				TaskState: "Disabled",
			},
			{
				TaskState: "Running",
			},
			{
				TaskState: "Running",
			},
			{
				TaskState: "Stopped",
			},
			{
				TaskState: "Stopped",
			},
			{
				TaskState: "Stopped",
			},
		},
		metrics: getTotalStatsMetricTypes(),
		expected: []string{
			"grafanalabs.snapstats.tasks.state.Disabled.count 1",
			"grafanalabs.snapstats.tasks.state.Running.count 2",
			"grafanalabs.snapstats.tasks.state.Stopped.count 3",
			"grafanalabs.snapstats.tasks.hitcount 0",
			"grafanalabs.snapstats.tasks.failedcount 0",
		},
	},
	{
		tasks: []Task{
			{
				HitCount:    10,
				FailedCount: 1,
			},
			{
				HitCount:    10,
				FailedCount: 1,
			},
			{
				HitCount:    10,
				FailedCount: 1,
			},
		},
		metrics: getTotalStatsMetricTypes(),
		expected: []string{
			"grafanalabs.snapstats.tasks.state.Disabled.count 0",
			"grafanalabs.snapstats.tasks.state.Running.count 0",
			"grafanalabs.snapstats.tasks.state.Stopped.count 0",
			"grafanalabs.snapstats.tasks.hitcount 30",
			"grafanalabs.snapstats.tasks.failedcount 3",
		},
	},
}

func TestCollector(t *testing.T) {
	Convey("When collecting total stats", t, func() {
		for _, c := range cases {
			metrics, err := collectTotalStats(c.metrics, c.tasks)

			So(err, ShouldBeNil)

			So(metrics, ShouldNotBeNil)
			So(len(metrics), ShouldEqual, len(c.expected))

			for i, metric := range metrics {
				So(format(&metric), ShouldEqual, c.expected[i])
			}
		}
	})
}

func format(m *plugin.Metric) string {
	return fmt.Sprintf("%v %v", strings.Join(m.Namespace.Strings(), "."), m.Data)
}
