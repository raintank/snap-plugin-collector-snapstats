package snapstats

import (
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/intelsdi-x/snap-plugin-utilities/logger"
	// "k8s.io/client-go/rest"
)

type Snapstats struct {
}

// CollectMetrics collects metrics for testing
func (n *Snapstats) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	LogDebug("request to collect metrics", "metric_count", len(mts))

	snapUrl, err := mts[0].Config.GetString("snap-url")
	if err != nil {
		LogError("The snapstats collector failed to fetch config value snap-url.", "error", err)
		return nil, err
	}

	client, err := NewClient(snapUrl, true)
	if err != nil {
		LogError("The snapstats collector failed to create Snap api client.", "error", err)
		return nil, err
	}

	tasks, err := client.GetTasks()

	metrics, err := collectTotalStats(mts, tasks)
	if err != nil {
		LogError("The snapstats collector failed to create TotalStats metrics.", "error", err)
		return nil, err
	}

	LogDebug("collecting metrics completed", "metric_count", len(metrics))
	return metrics, nil
}

func collectTotalStats(mts []plugin.Metric, tasks []Task) ([]plugin.Metric, error) {
	metrics := make([]plugin.Metric, 0)
	var totalCount int

	for _, mt := range mts {
		ns := mt.Namespace.Strings()
		totalCount = 0

		for _, t := range tasks {
			if ns[3] == "state" && t.TaskState == ns[4] {
				totalCount++
			} else if ns[3] == "hitcount" {
				totalCount += t.HitCount
			} else if ns[3] == "failedcount" {
				totalCount += t.FailedCount
			}
		}

		mt.Data = totalCount
		mt.Timestamp = time.Now()
		metrics = append(metrics, mt)
	}

	return metrics, nil
}

func getTotalStatsMetricTypes() []plugin.Metric {
	mts := []plugin.Metric{}

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("grafanalabs", "snapstats", "tasks", "state", "Disabled", "count"),
		Version:   1,
	})

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("grafanalabs", "snapstats", "tasks", "state", "Running", "count"),
		Version:   1,
	})

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("grafanalabs", "snapstats", "tasks", "state", "Stopped", "count"),
		Version:   1,
	})

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("grafanalabs", "snapstats", "tasks", "hitcount"),
		Version:   1,
	})

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("grafanalabs", "snapstats", "tasks", "failedcount"),
		Version:   1,
	})

	return mts
}

func (n *Snapstats) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}

	mts = append(mts, getTotalStatsMetricTypes()...)

	return mts, nil
}

func (f *Snapstats) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"grafanalabs", "snapstats"}, "snap-url", false, plugin.SetDefaultString("http://localhost:8181"))
	return *policy, nil
}
