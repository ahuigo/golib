package main
import (
        kafka "github.com/segmentio/kafka-go"
        "context"
)
func create_topic2(topic string){
    host1 := "192.168.0.1:9092"
    partition := 0
    conn, err := kafka.Dial("tcp", host1)
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    controller, err := conn.Controller()
    if err != nil {
        panic(err.Error())
    }
    var controllerConn *kafka.Conn
    controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
    if err != nil {
        panic(err.Error())
    }
    defer controllerConn.Close()


    topicConfigs := []kafka.TopicConfig{
        kafka.TopicConfig{
            Topic:             topic,
            NumPartitions:     1,
            ReplicationFactor: 1,
        },
    }

    err = controllerConn.CreateTopics(topicConfigs...)
    if err != nil {
        panic(err.Error())
    }
}
func create_topic(topic string){
    _, err := kafka.DialLeader(context.Background(), "tcp", "192.168.0.1:9092", topic, 0)
    if err != nil {
        panic(err.Error())
    }
}
func main(){
    create_topic("cadence-visibility-prod")
    create_topic("cadence-visibility-prod-dlq")
    // to create topics when auto.create.topics.enable='true'
}
