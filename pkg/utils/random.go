package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var (
	rands      = rand.New(rand.NewSource(time.Now().UnixNano()))
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ9876543210"
)

// RandomString 生成随机字符串，字符范围[a-zA-Z0-9]
func RandomString(n int) string {
	var sb bytes.Buffer
	var l = len(characters) - 1
	for i := 0; i < n; i++ {
		k := rands.Intn(l)
		_ = sb.WriteByte(characters[k])
	}

	return sb.String()
}

// RandomInt 随机生成数字[min, max]
func RandomInt(min, max int) int {
	return rands.Intn(max-min+1) + min
}

// RandomEmail 随机生成一个邮箱
func RandomEmail() string {
	nameLen := RandomInt(2, 12)
	nameStr := RandomString(nameLen)
	return fmt.Sprintf("%s@%s.%s", nameStr, RandomString(3), RandomString(3))
}
