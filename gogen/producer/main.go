package main

import (
	"belajar-rabbitmq/shared/driver"
	"belajar-rabbitmq/shared/infrastructure/messaging"
	"belajar-rabbitmq/shared/model/payload"
	"time"
)

func main() {

	pub := messaging.NewPublisher("amqp://guest:guest@localhost:5672/")

	err := pub.Publish("events.payment.ovo", payload.Payload{
		Data: "Hello World",
		Publisher: driver.ApplicationData{
			AppName:       "App1",
			AppInstanceID: "abc123",
			StartTime:     time.Now().Format("2006-01-02 15:04:05"),
		},
		TraceID: "aaaa",
	})
	if err != nil {
		panic(err.Error())
	}

}
