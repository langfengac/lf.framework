package lf

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

//默认des key
const defaultDesKey = "unshfbep"

//DES加密方法
func DESEncrypt(origDataStr, keyStr string) string {
	origData := []byte(origDataStr)
	key := []byte(keyStr)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//对明文先进行补码操作
	origData = pKCS5Padding(origData, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, key)
	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))
	//加密明文,加密后的数据放到数组中
	blockMode.CryptBlocks(crypted, origData)
	//将字节数组转换成字符串
	return base64.StdEncoding.EncodeToString(crypted)
}
func DESEncryptDefault(origDataStr string) string {
	return DESEncrypt(origDataStr, defaultDesKey)
}

//实现明文的补码
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(ciphertext, padtext...)
}

//解密
func DESDecrypt(data, keyStr string) string {
	key := []byte(keyStr)
	//倒叙执行一遍加密方法
	//将字符串转换成字节数组
	crypted, _ := base64.StdEncoding.DecodeString(data)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, key)
	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)
	//去补码
	origData = pKCS5UnPadding(origData)
	return string(origData)
}
func DESDecryptDefault(data string) string {
	return DESDecrypt(data, defaultDesKey)
}

//去除补码
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}
