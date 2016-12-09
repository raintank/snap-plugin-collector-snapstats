package main

import (
	// Import the snap plugin library
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	// Import our collector plugin implementation
	. "github.com/intelsdi-x/snap-plugin-utilities/logger"
	"github.com/raintank/snap-plugin-collector-snapstats/snapstats"
)

const (
	pluginName    = "snapstats"
	pluginVersion = 1
)

func main() {
	LogDebug("Starting snapstats collector")

	plugin.StartCollector(new(snapstats.Snapstats), pluginName, pluginVersion, plugin.ConcurrencyCount(1000))
}
