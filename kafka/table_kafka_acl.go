package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKafkaAcl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_acl",
		Description: "Get details of all the acls in your Kafka cluster.",
		List: &plugin.ListConfig{
			Hydrate: listAcls,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "resource_name",
					Require: plugin.Optional,
				},
				{
					Name:    "resource_type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Kafka acluration property. It specifies the unique identifier for the acluration setting.",
			},
			{
				Name:        "resource_name",
				Type:        proto.ColumnType_STRING,
				Description: "This field holds the name or identifier of the specific Kafka acluration resource. It refers to the individual entity for which the acluration properties are being defined or managed.",
			},
			{
				Name:        "resource_type",
				Type:        proto.ColumnType_INT,
				Description: "This field represents the type of Kafka acluration resource. It specifies the category or type of resource to which the acluration properties belong.",
				Transform:   transform.FromQual("resource_type"),
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "The current value of the Kafka acluration property. It holds the actual setting that dictates the behavior or characteristics of the Kafka system.",
			},
			{
				Name:        "read_only",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the acluration property is read-only. If set to true, the property cannot be modified at runtime.",
			},
			{
				Name:        "default",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the current value of the acluration property is the default value. If set to true, the property value is the default provided by Kafka.",
			},
			{
				Name:        "source",
				Type:        proto.ColumnType_JSON,
				Description: "A JSON representation of the source or origin of the acluration property. It may include information about where the property was defined or set, such as a acluration file or an environment variable.",
			},
			{
				Name:        "sensitive",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the acluration property contains sensitive information. If set to true, the value might include credentials or other confidential data.",
			},
			{
				Name:        "synonyms",
				Type:        proto.ColumnType_JSON,
				Description: "A JSON array containing synonyms or alternative names for the acluration property. This allows the same acluration setting to be referred to by multiple names.",
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

func listAcls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_acl.listAcls", "connection_error", err)
		return nil, err
	}
	var filter sarama.AclFilter
	//filter.ResourceName = types.String("mytopic1")
	//filter.ResourceType = 2
	//resource.Name = d.EqualsQualString("resource_name")
	//resource.Type = sarama.AclResourceType(d.EqualsQuals["resource_type"].GetInt64Value())
	acls, err := kafkaClient.Admin.ListAcls(filter)
	if err != nil {
		logger.Error("kafka_acl.listAcls", "api_error", err)
		return nil, err
	}

	for _, acl := range acls {
		d.StreamListItem(ctx, acl)
	}

	return nil, nil
}
