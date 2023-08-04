package main

import (
	"fmt"

	"github.com/IBM/sarama"
	// "github.com/turbot/steampipe-plugin-kafka/kafka"
	// "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		PluginFunc: kafka.Plugin})
// }

func main() {
	// Define the Kafka brokers' address.
	brokers := []string{"localhost:9092"}

	// Create a new cluster admin
	config := sarama.NewConfig()
	admin, err := sarama.NewClient(brokers, config)
	if err != nil {
		fmt.Printf("Failed to create cluster admin: %s\n", err.Error())
		return
	}
	var resource sarama.ConfigResource
	resource.Name = ""
	// Get the list of all brokers in the cluster.
	topicList, err := admin.Controller()
	if err != nil {
		fmt.Printf("Failed to create cluster admin: %s\n", err.Error())
		return
	}
	test, err := topicList.DescribeConfigs(&sarama.DescribeConfigsRequest{
		Resources: []*sarama.ConfigResource{&resource}})
	// Print the topic information
	//fmt.Println(topicList)
	for _, topic := range test.Resources {
		fmt.Println(topic.Type)

	}
}
