package models

import (
	"order/config"

	"github.com/streadway/amqp"
)

var MQ = Init_RabbitMQ(config.RabbitMQSetting.Url)

func Init_RabbitMQ(connString string) *amqp.Connection {
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	return conn
}
