package main

import (
	"log"
	"mq-server/service"
)

func main() {
	log.Println("[main] RabbitMQ is starting...")

	forever := make(chan bool)
	service.CreateOrder()
	<-forever
}
