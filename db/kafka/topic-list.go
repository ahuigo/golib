package main
import (
	"fmt"
 
	"github.com/Shopify/sarama"
)
 
func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
 
	//kafka end point
    brokers := []string{"192.168.0.1:9092"}
 
	//get broker
	cluster, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
 
	defer func() {
		if err := cluster.Close(); err != nil {
			panic(err)
		}
	}()
 
	//get all topic from cluster
	topics, _ := cluster.Topics()
	for index := range topics {
		fmt.Println(topics[index])
	}
}
