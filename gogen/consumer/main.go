package main

import (
	"belajar-rabbitmq/shared/infrastructure/messaging"
	"belajar-rabbitmq/shared/model/payload"
	"fmt"
)

func main() {

	subs := messaging.NewSubscriber("Q1")

	{
		topic := "events.payment.dana"
		subs.Handle(topic, func(payload payload.Payload, err error) {
			if err != nil {
				fmt.Printf("ERROR %v\n", err.Error())
				return
			}
			fmt.Printf("Recv %s : %v\n", topic, payload.Data)
		})
	}

	{
		topic := "events.payment.ovo"
		subs.Handle(topic, func(payload payload.Payload, err error) {
			if err != nil {
				fmt.Printf("ERROR %v\n", err.Error())
				return
			}
			fmt.Printf("Recv %s : %v\n", topic, payload.Data)
		})
	}

	{
		topic := "events.payment.*"
		subs.Handle(topic, func(payload payload.Payload, err error) {
			if err != nil {
				fmt.Printf("ERROR %v\n", err.Error())
				return
			}
			fmt.Printf("Recv %s : %v\n", topic, payload.Data)
		})
	}

	subs.Run("amqp://guest:guest@localhost:5672/")
}
