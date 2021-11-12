package rabbitmq

import (
	"github.com/langfengac/lf.framework/nlog"
	"github.com/streadway/amqp"
	"strconv"
)

//延迟队列
func NewDelayConsume(delayQueueName string, conn *amqp.Connection, ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	//辅助交换器和队列，默认创建
	exChangeName := "delay_exchange_for_" + delayQueueName
	queueName := "delay_queueName_for_" + delayQueueName
	// 声明交换器
	err := ch.ExchangeDeclare(exChangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	// 声明一个常规的队列, 其实这个也没必要声明,因为 exchange 会默认绑定一个队列
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	//声明延时队列队列，该队列中消息如果过期，就将消息发送到交换器上，交换器就分发消息到普通队列(这部分操作是在消息生产者上面)
	_, err = ch.QueueDeclare(delayQueueName, false, false, false, false,
		amqp.Table{
			// 当消息过期时把消息发送到 logs 这个 exchange
			"x-dead-letter-exchange": exChangeName,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(q.Name, "", exChangeName, false, nil)
	if err != nil {
		return nil, err
	}

	// 这里监听的是queueName
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

//发送延时队列消息
func PublishDelay(delayQueueName string, body string, expireSeconds int) bool {
	//连接rabbitmq
	conn := NewConnDefault()
	if conn == nil {
		return false
	}
	defer conn.Close()

	//打开通道
	ch, _ := conn.Channel()
	defer ch.Close()

	expiration := strconv.Itoa(expireSeconds * 1000) //过期时间  毫秒
	//生产消息
	err := ch.Publish("", delayQueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
		Expiration:  expiration, // 设置过期时间
	})
	if err != nil {
		nlog.Error(err, "写入消息队列错误 PublishDelay")
		return false
	}
	return true
}
