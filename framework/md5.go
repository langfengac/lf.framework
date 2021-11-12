package lf

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(items ...string) string {
	var text string
	for _, value := range items {
		text = text + value
	}
	h := md5.New()
	h.Write([]byte(text))
	sign := hex.EncodeToString(h.Sum(nil))
	return sign
}