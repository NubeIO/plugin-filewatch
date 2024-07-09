package main

import (
	"github.com/NubeIO/plugin-filewatch/watcher"
	"github.com/NubeIO/rxlib/libs/pluginlib"
)

func main() {
	factory := pluginlib.New("test")
	factory.AddPallet("watcher", watcher.New)
	factory.Register("test")
}
