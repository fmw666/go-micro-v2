package main

import "mq-server/service"

func main() {
	forever := make(chan bool)
	service.CreateOrder()
	<-forever
}
