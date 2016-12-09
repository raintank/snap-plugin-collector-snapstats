package snapstats

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/intelsdi-x/snap-plugin-utilities/logger"
	// "k8s.io/client-go/rest"
)

type Snapstats struct {
}

// CollectMetrics collects metrics for testing
func (n *Snapstats) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	LogDebug("request to collect metrics", "metric_count", len(mts))
	metrics := make([]plugin.Metric, 0)

	snapUrl, err := mts[0].Config.GetString("snap-url")
	if err != nil {
		LogError("failed to fetch config value snap-url.", "error", err)
		return nil, err
	}

	client, err := NewClient(snapUrl, true)
	if err != nil {
		LogError("failed to create Snap api client.", "error", err)
		return nil, err
	}

	client.GetTasks()
	// for _, t := range totalStats.Items {
	// 	podMetrics, _ := podCollector.Collect(mts, p)
	// 	metrics = append(metrics, podMetrics...)
	// }

	LogDebug("collecting metrics completed", "metric_count", len(metrics))
	return metrics, nil
}

func (n *Snapstats) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
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

	return mts, nil
}

func (f *Snapstats) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"grafanalabs", "snapstats"}, "snap-url", false, plugin.SetDefaultString("http://localhost:8181"))
	return *policy, nil
}
