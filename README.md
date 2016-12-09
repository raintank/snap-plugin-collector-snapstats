# snap collector plugin - snapstats

This plugin collects metrics from Snap itself. It is quite simple so far and just collects some counts for all tasks.

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [License](#license-and-authors)
4. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.5+](https://golang.org/dl/) (needed only for building)

### Installation

#### To build the plugin binary:
Fork https://github.com/raintank/snap-plugin-collector-snapstats
Clone repo into `$GOPATH/src/github.com/raintank/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-snapstats.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/`

Uses govendor to manage dependencies.

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/grafanalabs/snapstats/tasks/state/Disabled/count | Total number of disabled Snap tasks
/grafanalabs/snapstats/tasks/state/Running/count | Total number of Running Snap tasks
/grafanalabs/snapstats/tasks/state/Stopped/count | Total number of Stopped Snap tasks
/grafanalabs/snapstats/tasks/hitcount | Total sum of the all the snap tasks' hit counts
/grafanalabs/snapstats/tasks/failedcount | Total sum of the all the snap tasks' failed counts

### Examples

There is an [example task](examples/task.json) in the examples directory.

It has one config variable that is required - **snap-url**

For most users, this value should be `http://localhost:8181`. It is mandatory that the url starts with the http or https protocol.

### Roadmap

Hopefully, Intel will be releasing a more complete lib that will make this obsolete.

## License

This plugin is released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [@daniellee](https://github.com/daniellee/)