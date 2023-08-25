package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableKafkaTopicMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kafka_topic_message",
		Description: "Get all the messages from all the topics in your Kafka cluster.",
		List: &plugin.ListConfig{
			ParentHydrate: listTopics,
			Hydrate:       listTopicMessages,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "topic",
					Require: plugin.Optional,
				},
				{
					Name:    "offset",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "headers",
				Type:        proto.ColumnType_JSON,
				Description: "Headers can be used to store metadata or any additional information about the message. Each RecordHeader typically consists of a key-value pair.",
			},
			{
				Name:        "timestamp",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Represents the timestamp of when the individual message was produced or created.",
			},
			{
				Name:        "block_timestamp",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "This is the timestamp for the outer (compressed) block that contains this message. Kafka messages can be batched together and compressed for more efficient transfer and storage. This timestamp refers to the entire batch or block rather than the individual message.",
			},
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Description: "A byte slice representing the key of the message. Keys in Kafka can be used to determine which partition a message is sent to. This can be especially important when the order of messages matters, as only messages within a single partition are guaranteed to be in order.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "A byte slice representing the actual content or payload of the message. This is what consumers typically process or store.",
			},
			{
				Name:        "topic",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Kafka topic to which this message belongs.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_INT,
				Description: "An integer representing the partition number within the topic from which this message was consumed.",
			},
			{
				Name:        "offset",
				Type:        proto.ColumnType_INT,
				Description: "An integer representing the position or index of this message within its partition.",
			},
		},
	}
}

//// LIST FUNCTION

func listTopicMessages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	topic := h.Item.(*sarama.TopicMetadata)

	// check if the provided topic is matching with the parent hydrate
	if d.EqualsQualString("topic") != "" && d.EqualsQualString("topic") != topic.Name {
		return nil, nil
	}

	// Create client
	kafkaClient, err := getClient(ctx, d)
	if err != nil {
		logger.Error("kafka_topic_message.listTopicMessages", "connection_error", err)
		return nil, err
	}

	offSet := d.EqualsQuals["offset"].GetInt64Value()
	queryTime := time.Now()
	var wg sync.WaitGroup
	errorCh := make(chan error, len(topic.Partitions))
	for _, partition := range topic.Partitions {
		wg.Add(1)
		go getMsgDataAsync(ctx, partition, kafkaClient, &wg, topic.Name, errorCh, queryTime, d, offSet)

	}
	wg.Wait()
	close(errorCh)

	for err := range errorCh {
		// return the first error
		plugin.Logger(ctx).Error("kafka_topic_message.listTopicMessages", "channel_error", err)
		return nil, err
	}

	return nil, nil
}

func getMsgDataAsync(ctx context.Context, partition *sarama.PartitionMetadata, kafkaClient *kafkaClient, wg *sync.WaitGroup, topic string, errorCh chan error, queryTime time.Time, d *plugin.QueryData, offSet int64) {
	defer wg.Done()

	err := getMsgData(ctx, partition, kafkaClient, topic, queryTime, d, offSet)
	if err != nil {
		errorCh <- err
	}
}

func getMsgData(ctx context.Context, partition *sarama.PartitionMetadata, kafkaClient *kafkaClient, topic string, queryTime time.Time, d *plugin.QueryData, offSet int64) error {
	newConsumer, err := sarama.NewConsumer(kafkaClient.BootstrapServers, kafkaClient.Config)
	if err != nil {
		plugin.Logger(ctx).Error("kafka_topic_message.NewConsumer", "api_error", err)
		return err
	}

	// check if offset is provided by the user
	var startingOffset int64
	if offSet == 0 {
		startingOffset = sarama.OffsetOldest
	} else {
		startingOffset = offSet
	}

	consumer, err := newConsumer.ConsumePartition(topic, partition.ID, startingOffset)
	if err != nil {
		plugin.Logger(ctx).Error("kafka_topic_message.ConsumePartition", "api_error", err)
		return err
	}

	// Set the timer with default waiting period 2 secs
	idleDuration := time.Second * 2
	timeout := time.NewTimer(idleDuration)

	msgCount := 0
ConsumerLoop:
	for {
		timeout.Reset(idleDuration)

		// check if the number of messages consumed is matched the provided limit
		if d.QueryContext.Limit != nil && *d.QueryContext.Limit == int64(msgCount) {
			break ConsumerLoop
		}
		select {
		case msg := <-consumer.Messages():

			// check if the msg.Timestamp is less than the query execution time
			if msg.Timestamp.Before(queryTime) {
				d.StreamListItem(ctx, msg)
				msgCount++
			} else {
				break ConsumerLoop
			}
			continue
		case <-timeout.C:
			break ConsumerLoop
		}
	}

	return nil
}
