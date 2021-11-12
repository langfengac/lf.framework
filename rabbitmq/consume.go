package rabbitmq

import "github.com/streadway/amqp"

//消费
func NewConsume(queueName string, conn *amqp.Connection, ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	//声明队列
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	//消费消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
