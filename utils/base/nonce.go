package base

import (
	"math/rand"
	"time"
)

// NonceStr 随机字符
func NonceStr(length int) string {
	var nonceStr string
	var char = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		r := rand.Uint64()
		key := r % uint64(58)
		str := char[key]
		nonceStr += string(str)
	}
	return nonceStr
}

// NonceNumb 随机数字
func NonceNumb(length int) string {
	var nonceStr string
	var char = "0123456789"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		r := rand.Uint64()
		key := r % uint64(10)
		str := char[key]
		nonceStr += string(str)
	}
	return nonceStr
}
