package main

import (
	"github.com/turbot/steampipe-plugin-kafka/kafka"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//	"github.com/IBM/sarama"
//"github.com/turbot/steampipe-plugin-kafka/kafka"
//"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: kafka.Plugin})
}

// func main() {
// 	// Define the Kafka brokers' address.
// 	brokers := []string{"localhost:9092"}

// 	// Create a new cluster admin
// 	config := sarama.NewConfig()
// 	admin, err := sarama.NewClusterAdmin(brokers, config)
// 	if err != nil {
// 		fmt.Printf("Failed to create cluster admin: %s\n", err.Error())
// 		return
// 	}
// 	var resource sarama.AclFilter
// 	resource.Version = 29
// 	resource.ResourceType = sarama.AclResourceTopic
// 	resource.ResourcePatternTypeFilter = sarama.AclPatternAny
// 	resource.PermissionType = sarama.AclPermissionAny
// 	resource.Operation = sarama.AclOperationAny

// 	// Get the list of all brokers in the cluster.
// 	topicList, err := admin.ListAcls(resource)
// 	if err != nil {
// 		fmt.Printf("Failed to create cluster admin: %s\n", err.Error())
// 		return
// 	}
// 	// test, err := topicList.DescribeConfigs(&sarama.DescribeConfigsRequest{
// 	// 	Resources: []*sarama.ConfigResource{&resource}})
// 	// Print the topic information
// 	//fmt.Println(topicList)
// 	for _, topic := range topicList {
// 		for _, acl := range topic.Acls {
// 			fmt.Println(acl.Principal)
// 		}
// 	}
// }
