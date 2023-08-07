package main

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	// "github.com/turbot/steampipe-plugin-kafka/kafka"
	// "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		PluginFunc: kafka.Plugin})
// }

func main() {
	seeds := []string{"localhost:9092"}

	var adminClient *kadm.Client
	{
		client, err := kgo.NewClient(
			kgo.SeedBrokers(seeds...),

			// Do not try to send requests newer than 2.4.0 to avoid breaking changes in the request struct.
			// Sometimes there are breaking changes for newer versions where more properties are required to set.
			//kgo.MaxVersions(kversion.V2_4_0()),
		)
		if err != nil {
			panic(err)
		}
		defer client.Close()

		adminClient = kadm.NewClient(client)
	}
	b := &kadm.ACLBuilder{}
	c := b.Topics()
	topicList, err := adminClient.DescribeACLs(context.Background(), c)
	if err != nil {
		fmt.Printf("Failed to create cluster admin: %s\n", err.Error())
		return
	}
	// test, err := topicList.DescribeConfigs(&sarama.DescribeConfigsRequest{
	// 	Resources: []*sarama.ConfigResource{&resource}})
	// Print the topic information
	fmt.Println(topicList)
	for _, topic := range topicList {
		fmt.Println(topic.Principal)

	}
}
