package main

import (
	"github.com/turbot/steampipe-plugin-kafka/kafka"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: kafka.Plugin})
}
