package kafka

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-kafka"

// Plugin creates this (kafka) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"kafka_broker":         tableKafkaBroker(ctx),
			"kafka_config":         tableKafkaConfig(ctx),
			"kafka_consumer_group": tableKafkaConsumerGroup(ctx),
			"kafka_topic":          tableKafkaTopic(ctx),
		},
	}

	return p
}
