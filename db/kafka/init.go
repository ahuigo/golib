package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Version = sarama.V1_0_0_0
	//kafka end point
    brokers := []string{"192.168.0.1:9092", "192.168.0.2:9092"}

	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		fmt.Println("error is", err)
		return
	}
	detail := sarama.TopicDetail{NumPartitions: 2, ReplicationFactor: 1}

	err = admin.CreateTopic("cadence-visibility-prod", &detail, true)
	if err != nil {
		fmt.Println("error is", err)
	}
	err = admin.CreateTopic("cadence-visibility-prod-dlq", &detail, true)
	if err != nil {
		fmt.Println("error is", err)
	}
}
