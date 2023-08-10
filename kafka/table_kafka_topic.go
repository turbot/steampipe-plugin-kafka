package kafka

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKafkaTopic(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_topic",
		Description: "Get details of all the topics in your Kafka cluster.",
		List: &plugin.ListConfig{
			Hydrate: listTopics,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "version",
				Type:        proto.ColumnType_INT,
				Description: "Defines the protocol version to use for encode and decode.",
			},
			{
				Name:        "err",
				Type:        proto.ColumnType_JSON,
				Description: "Contains the topic error, or 0 if there was no error.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Contains the topic name.",
			},
			{
				Name:        "is_internal",
				Type:        proto.ColumnType_BOOL,
				Description: "Contains a true if the topic is internal.",
			},
			{
				Name:        "partitions",
				Type:        proto.ColumnType_JSON,
				Description: "Contains each partition in the topic.",
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

func listTopics(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_topic.listTopics", "connection_error", err)
		return nil, err
	}

	// fetch topic names
	topicName := d.EqualsQualString("name")
	allTopics := []string{}
	if topicName == "" {
		topics, err := kafkaClient.Client.Topics()
		if err != nil {
			logger.Error("kafka_topic.Topics", "api_error", err)
			return nil, err
		}
		allTopics = append(allTopics, topics...)
	} else {
		allTopics = append(allTopics, topicName)
	}

	topicsMetadata, err := kafkaClient.Admin.DescribeTopics(allTopics)
	if err != nil {
		logger.Error("kafka_topic.DescribeTopics", "api_error", err)
		return nil, err
	}

	for _, topic := range topicsMetadata {
		d.StreamListItem(ctx, topic)
	}

	return nil, nil
}
