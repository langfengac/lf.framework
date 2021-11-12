package rabbitmq

import (
	"github.com/langfengac/lf.framework/appconfig"
	"github.com/langfengac/lf.framework/nlog"
	"github.com/streadway/amqp"
)

func NewConn(confKey string) *amqp.Connection {
	url := getConn(confKey)
	conn, err := amqp.Dial("amqp://" + url)
	if err != nil {
		nlog.Error(err, "消息队列连接错误 NewConn")
		return nil
	}
	return conn
}
func NewConnDefault() *amqp.Connection {
	return NewConn("default")
}

func getConn(confKey string) string {
	c := appconfig.NewInitAppConfig()
	s := c.ReadString("rabbitmq", confKey, "")
	//if appconfig.IsRelease() {
	//	//解密
	//	s = lf.DESDecryptDefault(s)
	//}
	return s
}