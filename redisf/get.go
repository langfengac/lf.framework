package redisf

func Get(redisKey string) string {
	client := DefaultClient(0)
	defer client.Close()
	str, _ := client.Get(redisKey).Result()
	return str
}
