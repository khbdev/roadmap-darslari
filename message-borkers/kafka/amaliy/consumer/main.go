package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "user.created"
	groupID := "users-service"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset, 
	})
	defer r.Close()

	fmt.Println("Listening... topic =", topic, "group =", groupID)


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-stop:
			fmt.Println("\n Stopped")
			return
		default:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			msg, err := r.ReadMessage(ctx)
			cancel()

			if err != nil {
			
				continue
			}

			fmt.Printf("Got message | partition=%d offset=%d key=%s value=%s\n",
				msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}
