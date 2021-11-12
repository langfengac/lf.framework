package main

import (
	"errors"
	"fmt"
	"github.com/langfengac/lf.framework/appconfig"
	lf "github.com/langfengac/lf.framework/framework"
	"github.com/langfengac/lf.framework/nlog"
	"github.com/langfengac/lf.framework/rabbitmq"
	"math/rand"
)

func main() {
	//test()
	//mqc()
	//mqp()
	//mqc_delay()
	//mqp_delay()
	des()
}
func test() {
	r := rand.Intn(9999)
	fmt.Println(r)

	c := appconfig.NewInitAppConfig()
	s := c.ReadString("nlog", "dir_path", "")
	fmt.Println(s)

	fmt.Println("解密:", c.ReadString("t", "test", ""))

	nlog.Info("日志text")
	nlog.Error(errors.New("我是一个错误"), "自定义描述")

	//cc := "unifygoods:cu9pRR4nb4942c4d@tcp(10.31.239.91:3306)/nd99_unifygoods?charset=utf8"
	//cc := "10.31.239.29:6379|D4Ubu7g3eJy"
	cc := "zhifu:9nXtAVHqXhjU9FvZ@10.31.239.29:5672/"
	ee := lf.DESEncryptDefault(cc)
	fmt.Println(cc)
	fmt.Println(ee)

	tStr := "20210208110159"
	t, _ := lf.Str14ToTime(tStr)
	tStr = lf.TimeToStr14(t)
	fmt.Println(t, tStr)
}
func mqp() {
	rabbitmq.Publish("go_test1", "i am a message")
}
func mqc() {
	fmt.Println("消息队列任务启动")

	//连接rabbitmq
	conn := rabbitmq.NewConnDefault()
	defer conn.Close()

	//打开通道
	ch, _ := conn.Channel()
	defer ch.Close()

	//消费消息
	msgs, _ := rabbitmq.NewConsume("go_test1", conn, ch)

	go func() {
		for msg := range msgs {
			//因为存入的是byt[] 所以取出后转成字符串打印	string(msg.Body)
			s := string(msg.Body)
			fmt.Printf("收到队列消息：%v 数据长度：%d \r\n", s, len(s))
			//fmt.Println(i, "收到队列消息 数据长度：", len(msg.Body))
		}
	}()
}
func mqp_delay() {
	rabbitmq.PublishDelay("test_delay_3", "i am a message", 3)
}
func mqc_delay() {
	fmt.Println("延迟消息队列任务启动")

	//连接rabbitmq
	conn := rabbitmq.NewConnDefault()
	defer conn.Close()

	//打开通道
	ch, _ := conn.Channel()
	defer ch.Close()

	//消费消息
	msgs, _ := rabbitmq.NewDelayConsume("test_delay_3", conn, ch)

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			//因为存入的是byt[] 所以取出后转成字符串打印	string(msg.Body)
			s := string(msg.Body)
			fmt.Printf("收到队列消息：%v 数据长度：%d \r\n", s, len(s))
			//fmt.Println(i, "收到队列消息 数据长度：", len(msg.Body))
		}
	}()

	<-forever
}
func des() {
	fmt.Println("【解密】")
	t := "bJ2bLWp2sTqp2zaaT221zGZ0nQ7Yibm5BXABKJ4ZVRvCnFM63U4SyQ7V8wbz+4DCp9sMI9X/YuAPMj7OV5whCIvs69qsfQMVPJoCdJT8S8Q="
	d := lf.DESDecryptDefault(t)
	fmt.Println(d)

	t = "/9UIHZuOvgCXT3JFggieyY3ldgvawVTpSMPSNVwacD4="
	d = lf.DESDecryptDefault(t)
	fmt.Println(d)

	fmt.Println("【加密】")
	m1 := "guest:guest@172.24.132.232:5672/"
	m2 := lf.DESEncryptDefault(m1)
	fmt.Println("结果=", m2)

	//海外redis
	m1 = "10.55.1.85:6379|9a9ee8afffab"
	m2 = lf.DESEncryptDefault(m1)
	fmt.Println("海外redis=", m2)

	//海外mysql
	m1 = "hwndunifygoods:bwpptbvy9rZVGmq7@tcp(10.55.1.78:3306)/nd99_unifygoods?charset=utf8"
	m2 = lf.DESEncryptDefault(m1)
	fmt.Println("海外mysql=", m2)

	//日志追踪logs
	m1 = "logs:ZfyrD3568SSYKGP9@tcp(10.31.239.91:3306)/logs?charset=utf8"
	m2 = lf.DESEncryptDefault(m1)
	fmt.Println("日志追踪logs=", m2)

	//mq
	m1 = "zhifu:9nXtAVHqXhjU9FvZ@10.31.239.29:5672/"
	m2 = lf.DESEncryptDefault(m1)
	fmt.Println("mq=", m2)

	//es
	m1 = "http://10.31.239.100:9330"
	m2 = lf.DESEncryptDefault(m1)
	fmt.Println("es=", m2)
}
