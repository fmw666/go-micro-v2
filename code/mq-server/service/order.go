package service

import (
	"encoding/json"
	"log"
	"mq-server/models"
)

// 从 RabbitMQ 中接收信息，写入数据库
func CreateOrder() {
	ch, err := models.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("order", true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 处于一个监听状态，一致监听我们的生成端的生产，所以这里我们要阻塞主进程
	go func() {
		for d := range msgs {
			var t models.Order
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			models.DB.Create(&t)
			log.Println("[Order] Create Order:", t)
			_ = d.Ack(false)
		}
	}()
}
