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
	groupID := "users-service" // consumer group nomi

	// Reader = consumer. GroupID bersang -> consumer group ishlaydi.
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset, // group yangi bo'lsa, boshidan o'qisin
	})
	defer r.Close()

	fmt.Println("👂 Listening... topic =", topic, "group =", groupID)

	// Ctrl+C bo'lsa chiqib ketish uchun
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-stop:
			fmt.Println("\n🛑 Stopped")
			return
		default:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			msg, err := r.ReadMessage(ctx)
			cancel()

			if err != nil {
				// timeout bo‘lishi mumkin, davom etaveradi
				continue
			}

			fmt.Printf("✅ Got message | partition=%d offset=%d key=%s value=%s\n",
				msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}
