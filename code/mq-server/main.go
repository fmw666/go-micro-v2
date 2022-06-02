package main

import (
	"mq-server/config"
	"mq-server/models"
	"mq-server/service"
)

func main() {
	// 初始化 mysal
	models.Database(config.DatabaseSetting.Url)
	// 初始化 mq
	models.RabbitMQ(config.RabbitMQSetting.Url)

	forever := make(chan bool)
	service.CreateOrder()
	<-forever
}
