package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type UserCreated struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	At    string `json:"at"`
}

func main() {
	topic := "user.created"

	// Writer = producer. U topic'ga message yozadi.
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // qaysi partition'ga tushishini "tarqatadi"
	})
	defer w.Close()

	event := UserCreated{
		ID:    1,
		Name:  "Azizbek",
		Email: "azizbek@example.com",
		At:    time.Now().Format(time.RFC3339),
	}

	value, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	// Key bo'lsa: bir xil key -> doim bir xil partition (tartib saqlanadi)
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("user:%d", event.ID)),
		Value: value,
		Time:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := w.WriteMessages(ctx, msg); err != nil {
		panic(err)
	}

	fmt.Println("✅ Sent to", topic, "=>", string(value))
}
