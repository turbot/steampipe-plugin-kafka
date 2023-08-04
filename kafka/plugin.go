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
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"resource not found"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"kafka_topic":          tableKafkaTopic(ctx),
			"kafka_consumer_group": tableKafkaConsumerGroup(ctx),
			// "dockerhub_token":      tableDockerHubToken(ctx),
			// "dockerhub_user":       tableDockerHubUser(ctx),
		},
	}

	return p
}
