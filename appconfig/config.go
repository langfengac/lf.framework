package appconfig

import (
	"bufio"
	"fmt"
	"io"
	lf "lf.framework/framework"
	"os"
	"strconv"
	"strings"
)

const middle = "========="
const defaultConfig = "app.config"
const defaultClass = "default"

type Config struct {
	Mymap  map[string]string
	strcet string
}

func NewInitConfig(path string) *Config {
	c := new(Config)
	c.InitConfig(path)
	return c
}
func NewInitAppConfig() *Config {
	return NewInitConfig(defaultConfig)
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		//defer func() {
		//	if err := recover(); err != nil {
		//		fmt.Println("objToEsAllocation错误", err)
		//	}
		//}()

		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		//key
		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		//value
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		//判断是否需要解密
		if frist[0:1] == "*" {
			frist = frist[1:]
			second = decrypt(second)
			fmt.Println("second", second)
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}
func decrypt(s string) (result string) {
	defer func() {
		if err := recover(); err != nil {
			//如果错误返回空字符串
			result = ""
		}
	}()

	result = lf.DESDecryptDefault(s)
	return
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
func (c Config) ReadString(node, key, defaultValue string) string {
	s := c.Read(node, key)
	if s == "" {
		return defaultValue
	}
	return s
}
func (c Config) ReadInt(node, key string, defaultValue int) int {
	s := c.Read(node, key)
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return v
}
func (c Config) ReadFloat64(node, key string, defaultValue float64) float64 {
	s := c.Read(node, key)
	v, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return defaultValue
	}
	return v
}
func (c Config) ReadBool(node, key string, defaultValue bool) bool {
	s := c.Read(node, key)
	v := strings.ToLower(s) == "true"
	return v
}

func (c Config) ReadStringDefault(key, defaultValue string) string {
	return c.ReadString(defaultClass, key, defaultValue)
}
func (c Config) ReadIntDefault(key string, defaultValue int) int {
	return c.ReadInt(defaultClass, key, defaultValue)
}
func (c Config) ReadFloat64Default(key string, defaultValue float64) float64 {
	return c.ReadFloat64(defaultClass, key, defaultValue)
}
func (c Config) ReadBoolDefault(key string, defaultValue bool) bool {
	return c.ReadBool(defaultClass, key, defaultValue)
}

//是否发布 通用
func IsRelease() bool {
	return Release() > 0
}
func Release() int {
	c := NewInitAppConfig()
	return c.ReadIntDefault("release", 0)
}
