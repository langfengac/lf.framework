package rabbitmq

import (
	"github.com/streadway/amqp"
	"lf.framework/nlog"
)

//生产消息
func Publish(queueName string, body string) bool {
	//连接rabbitmq
	conn := NewConnDefault()
	if conn == nil {
		return false
	}
	defer conn.Close()

	//打开通道
	ch, _ := conn.Channel()
	defer ch.Close()

	//声明队列
	q, _ := ch.QueueDeclare(queueName, false, false, false, false, nil)

	//生产消息
	err := ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		nlog.Error(err, "写入消息队列错误 Publish")
		return false
	}
	return true
}

