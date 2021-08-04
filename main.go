package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)


type EmployeeEvent struct {
	Name string`json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()
	r.POST("/employee", func(c *gin.Context) {
		var employee EmployeeEvent
		c.BindJSON(&employee)

		fmt.Printf("EMPLOYEE to store: %v\n", employee)

		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "127.0.0.1:9092"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created Producer %v\n", p)
		defer p.Close()
		b, err := json.Marshal(employee)
		if err != nil {
			fmt.Printf("Error: %s", err)

		}
		deliveryChan := make(chan kafka.Event)
		topic := "trabajadores_test"

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(string(b)),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
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
