package kafka

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type kafkaConfig struct {
	BootstrapServers *[]string `cty:"bootstrap_servers"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"bootstrap_servers": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{
			Type: schema.TypeString,
		},
	},
}

func ConfigInstance() interface{} {
	return &kafkaConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) kafkaConfig {
	if connection == nil || connection.Config == nil {
		return kafkaConfig{}
	}
	config, _ := connection.Config.(kafkaConfig)
	return config
}
