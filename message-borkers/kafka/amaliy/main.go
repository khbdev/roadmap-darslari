package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	bootstrap := "localhost:9092"
	topic := "user.created"


	conn, err := kafka.Dial("tcp", bootstrap)
	if err != nil {
		panic(err)
	}
	defer conn.Close()


	controller, err := conn.Controller()
	if err != nil {
		panic(err)
	}

	
	ctrlAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)
	adminConn, err := kafka.Dial("tcp", ctrlAddr)
	if err != nil {
		panic(err)
	}
	defer adminConn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = adminConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(" Topic created:", topic)
	_ = ctx
}
