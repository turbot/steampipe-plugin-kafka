package kafka

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKafkaBroker(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_broker",
		Description: "Get details of all the brokers in your Kafka cluster.",
		List: &plugin.ListConfig{
			Hydrate: listBrokers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "ID returns the broker ID retrieved from Kafka's metadata, or -1 if that is not known.",
			},
			{
				Name:        "addr",
				Type:        proto.ColumnType_STRING,
				Description: "Addr returns the broker address as either retrieved from Kafka's metadata or passed to NewBroker.",
			},
			{
				Name:        "connected",
				Type:        proto.ColumnType_BOOL,
				Description: "Connected returns true if the broker is connected and false otherwise.",
			},
			{
				Name:        "rack",
				Type:        proto.ColumnType_STRING,
				Description: "Rack returns the broker's rack as retrieved from Kafka's metadata or the empty string if it is not known. The returned value corresponds to the broker's broker.rack configuration setting. Requires protocol version to be at least v0.10.0.0.",
			},
			{
				Name:        "response_size",
				Type:        proto.ColumnType_INT,
				Description: "The broker response size.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		},
	}
}

type Broker struct {
	Id           int32
	Addr         string
	Connected    bool
	Rack         string
	ResponseSize int
}

//// LIST FUNCTION

func listBrokers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_broker.listBrokers", "connection_error", err)
		return nil, err
	}

	brokers, _, err := kafkaClient.Admin.DescribeCluster()
	if err != nil {
		logger.Error("kafka_broker.listBrokers", "api_error", err)
		return nil, err
	}

	for _, broker := range brokers {
		connected, _ := broker.Connected()
		d.StreamListItem(ctx, Broker{broker.ID(), broker.Addr(), connected, broker.Rack(), broker.ResponseSize()})
	}

	return nil, nil
}
