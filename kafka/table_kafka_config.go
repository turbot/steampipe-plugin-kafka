package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKafkaConfig(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_config",
		Description: "Get details of all the configs of topics and brokers in your Kafka cluster.",
		List: &plugin.ListConfig{
			Hydrate: listConfigs,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "resource_name",
					Require: plugin.Required,
				},
				{
					Name:    "resource_type",
					Require: plugin.Required,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Kafka configuration property. It specifies the unique identifier for the configuration setting.",
			},
			{
				Name:        "resource_name",
				Type:        proto.ColumnType_STRING,
				Description: "This field holds the name or identifier of the specific Kafka configuration resource. It refers to the individual entity for which the configuration properties are being defined or managed.",
				Transform:   transform.FromQual("resource_name"),
			},
			{
				Name:        "resource_type",
				Type:        proto.ColumnType_INT,
				Description: "This field represents the type of Kafka configuration resource. It specifies the category or type of resource to which the configuration properties belong.",
				Transform:   transform.FromQual("resource_type"),
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "The current value of the Kafka configuration property. It holds the actual setting that dictates the behavior or characteristics of the Kafka system.",
			},
			{
				Name:        "read_only",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the configuration property is read-only. If set to true, the property cannot be modified at runtime.",
			},
			{
				Name:        "default",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the current value of the configuration property is the default value. If set to true, the property value is the default provided by Kafka.",
			},
			{
				Name:        "source",
				Type:        proto.ColumnType_JSON,
				Description: "A JSON representation of the source or origin of the configuration property. It may include information about where the property was defined or set, such as a configuration file or an environment variable.",
			},
			{
				Name:        "sensitive",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the configuration property contains sensitive information. If set to true, the value might include credentials or other confidential data.",
			},
			{
				Name:        "synonyms",
				Type:        proto.ColumnType_JSON,
				Description: "A JSON array containing synonyms or alternative names for the configuration property. This allows the same configuration setting to be referred to by multiple names.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listConfigs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_config.listConfigs", "connection_error", err)
		return nil, err
	}
	var resource sarama.ConfigResource
	resource.Name = d.EqualsQualString("resource_name")
	resource.Type = sarama.ConfigResourceType(d.EqualsQuals["resource_type"].GetInt64Value())
	configs, err := kafkaClient.Admin.DescribeConfig(resource)
	if err != nil {
		logger.Error("kafka_config.listConfigs", "api_error", err)
		return nil, err
	}

	for _, config := range configs {
		d.StreamListItem(ctx, config)
	}

	return nil, nil
}
