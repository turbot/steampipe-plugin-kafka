package kafka

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type kafkaClient struct {
	Admin            sarama.ClusterAdmin
	Client           sarama.Client
	BootstrapServers []string
	Config           *sarama.Config
}

func getClient(ctx context.Context, d *plugin.QueryData) (*kafkaClient, error) {
	conn, err := GetNewClientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*kafkaClient), nil
}

var GetNewClientCached = plugin.HydrateFunc(GetNewClientUncached).Memoize()

func GetNewClientUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	var bootstrapServers []string
	kafkaConfig := GetConfig(d.Connection)
	if kafkaConfig.BootstrapServers != nil {
		bootstrapServers = *kafkaConfig.BootstrapServers
	}

	// Error if the minimum config is not set
	if len(bootstrapServers) == 0 {
		return nil, errors.New("bootstrap_servers must be configured")
	}

	config := sarama.NewConfig()
	client, err := sarama.NewClient(bootstrapServers, config)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdmin(bootstrapServers, config)
	if err != nil {
		return nil, err
	}

	kafkaClient := &kafkaClient{
		Admin:            admin,
		Client:           client,
		BootstrapServers: bootstrapServers,
		Config:           config,
	}

	return kafkaClient, nil
}
