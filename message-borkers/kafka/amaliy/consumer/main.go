package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)



func main(){
	brokers := []string{"localhost:9092"}
	topic := "user.created"
	groupID := "commitTEST"	

		r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,

		CommitInterval: 0, // ✅ AUTO COMMIT OFF (manual commit)
	})
	defer r.Close()
		ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
		go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Println("stopping...")
		cancel()
	}()
		log.Println("MANUAL commit consumer started...")

		for {
				m, err := r.FetchMessage(ctx)
		if err != nil {
			// ctx cancel bo'lsa chiqib ketamiz
			if ctx.Err() != nil {
				break
			}
			log.Fatal("fetch error:", err)
		}

		fmt.Printf("READ  partition=%d offset=%d key=%s value=%s\n",
			m.Partition, m.Offset, string(m.Key), string(m.Value))
	if err := process(m); err != nil {
			log.Println("PROCESS FAIL -> commit YO'Q:", err)
			continue
		}
			if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("commit error:", err)
		}

		log.Println("COMMITTED offset:", m.Offset)

		}
			log.Println("bye")
}

func process(m kafka.Message) error {
	// demo uchun ozgina kutamiz
	time.Sleep(200 * time.Millisecond)

	// test qilish uchun shunaqa qilsang bo'ladi:
	// agar message value ichida "fail" bo'lsa error qaytar
	if string(m.Value) == "fail" {
		return fmt.Errorf("simulated fail")
	}
	return nil
}