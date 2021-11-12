package redisf

import (
	"github.com/go-redis/redis"
	"github.com/langfengac/lf.framework/appconfig"
	"strings"
)

func NewClient(addrConfKey string, db int) *redis.Client {
	s := getAddrs(addrConfKey)
	arr := strings.Split(s, "|")
	if len(arr) < 2 {
		return nil
	}
	address := arr[0]
	pwd := arr[1]
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: pwd,
		DB:       db,
	})
	return client
}

func DefaultClient(db int) *redis.Client {
	return NewClient("addrs", db)
}

func getAddrs(addrConfKey string) string {
	c := appconfig.NewInitAppConfig()
	s := c.ReadString("redis", addrConfKey, "")
	//if appconfig.IsRelease() {
	//	//解密
	//	s = lf.DESDecryptDefault(s)
	//}
	return s
}
