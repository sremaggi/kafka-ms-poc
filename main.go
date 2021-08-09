package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"kafka-producer/src/models"
)





func main() {
	r := gin.Default()
	r.POST("/:topic", func(c *gin.Context) {
		var event models.Request
		c.BindJSON(&event)


		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": event.Brokers})
		if err != nil {
			fmt.Printf("Producer error: %v\n", err)
		}
		fmt.Printf("Created Producer %v\n", p)
		defer p.Close()

		msgJSON, err := json.Marshal(event.ProducerMessage.Message)
		if err != nil {
			fmt.Printf("Error: %s", err)

		}
		deliveryChan := make(chan kafka.Event)
		fmt.Println("MESSAGGE::: ",string(msgJSON))
		topic := c.Param("topic")
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key: []byte(event.ProducerMessage.Key),
			Value:          []byte(string(msgJSON)),
			Headers:        []kafka.Header{kafka.Header{
				Value: []byte(event.ProducerMessage.Headers.Value),
				Key: event.ProducerMessage.Headers.Key,
			}},
		}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		close(deliveryChan)
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
