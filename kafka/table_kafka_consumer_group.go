package kafka

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKafkaConsumerGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_consumer_group",
		Description: "Get details of all the consumer groups in your Kafka cluster.",
		List: &plugin.ListConfig{
			Hydrate: listConsumerGroups,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "group_id",
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
				Description: "Contains the describe error as the KError type.",
			},
			{
				Name:        "error_code",
				Type:        proto.ColumnType_INT,
				Description: "Contains the describe error, or 0 if there was no error.",
			},
			{
				Name:        "group_id",
				Type:        proto.ColumnType_STRING,
				Description: "Contains the group ID string.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Contains the group state string, or the empty string.",
			},
			{
				Name:        "protocol_type",
				Type:        proto.ColumnType_STRING,
				Description: "Contains the group protocol type, or the empty string.",
			},
			{
				Name:        "protocol",
				Type:        proto.ColumnType_STRING,
				Description: "Contains the group protocol data, or the empty string.",
			},
			{
				Name:        "members",
				Type:        proto.ColumnType_JSON,
				Description: "Contains the group members.",
			},
			{
				Name:        "authorized_operations",
				Type:        proto.ColumnType_INT,
				Description: "Contains a 32-bit bit field to represent authorized operations for this group.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupId"),
			},
		},
	}
}

//// LIST FUNCTION

func listConsumerGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_consumer_group.listConsumerGroups", "connection_error", err)
		return nil, err
	}

	// fetch group ids
	gid := d.EqualsQualString("group_id")
	var groupIds []string
	if gid == "" {
		groups, err := kafkaClient.Admin.ListConsumerGroups()
		if err != nil {
			logger.Error("kafka_topic.listConsumerGroups", "api_error", err)
			return nil, err
		}
		for group := range groups {
			groupIds = append(groupIds, group)
		}
	} else {
		groupIds = append(groupIds, gid)
	}

	groupDescriptions, err := kafkaClient.Admin.DescribeConsumerGroups(groupIds)
	if err != nil {
		logger.Error("kafka_topic.DescribeConsumerGroups", "api_error", err)
		return nil, err
	}

	for _, groupDescription := range groupDescriptions {
		d.StreamListItem(ctx, groupDescription)
	}

	return nil, nil
}
