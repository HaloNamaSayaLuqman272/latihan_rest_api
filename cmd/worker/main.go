package main

import (
	"context"
	"encoding/json"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/events"
	"latihan_rest_api/pkg/rabbitmq"
	"latihan_rest_api/pkg/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mq, err := rabbitmq.New(utils.GetConfigurance(constant.RABBITMQ_URL))
	if err != nil {
		log.Fatalf("failed to connect rabbitmq: %v", err)
	}
	defer mq.Close()

	log.Println("worker started, waiting for messages")

	err = mq.Consume(ctx, "notification.user_registered", events.UserRegistered,
		func(ctx context.Context, body []byte) error {
			var ev events.UserRegisteredEvent
			if err := json.Unmarshal(body, &ev); err != nil {
				return err
			}

			log.Printf("send email wellcome to %s (%s)", ev.Username, ev.Email)
			return nil
		})

	if err != nil {
		log.Fatalf("consumer stopped: %v\n", err)
	}
}
