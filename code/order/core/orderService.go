package core

import (
	"context"
	"encoding/json"
	"order/models"
	"order/pkg/e"
	"order/service"

	"github.com/streadway/amqp"
)

// 微服务调用的 GetOrderList 方法
func (*OrderService) GetOrderList(ctx context.Context, req *service.OrderListRequest, resp *service.OrderListResponse) error {
	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if req.UserId != 0 {
		orders = orders.Where("user_id = ?", req.UserId)
	}
	err := orders.Count(&count).Offset(int(req.Offset)).Limit(int(req.Limit)).Order("id desc").Find(&data).Error
	if err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return nil
	}
	// 生成响应
	resp.Code = e.SUCCESS
	resp.Data = buildOrderList(data)
	resp.PageInfo = &service.PageInfo{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  uint32(count),
	}
	return nil
}

// 微服务调用的 CreateOrder 方法
func (*OrderService) CreateOrder(ctx context.Context, req *service.OrderCreateRequest, resp *service.OrderCreateResponse) error {
	// 将信息生产，放到 rabbitmq 队列中
	ch, err := models.MQ.Channel()
	if err != nil {
		resp.Code = e.ERROR_MQ_BASE
		resp.Message = e.GetMsg(e.ERROR_MQ_BASE)
		return nil
	}
	q, _ := ch.QueueDeclare("order_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		resp.Code = e.ERROR_MQ_PUBLISH
		resp.Message = e.GetMsg(e.ERROR_MQ_PUBLISH)
		return nil
	}
	// 将创建 order 任务交给 rabbitmq 处理，这里不需要返回创建的 order 信息
	resp.Code = e.SUCCESS
	return nil
}
