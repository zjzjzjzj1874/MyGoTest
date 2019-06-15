package util

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

// 生成传入长度的随机字符串(可包括数字,大小写字母)
func GetRandomString(length int, typeString string) string {
	var num = "0123456789"
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := bytes.Buffer{}
	if strings.Contains(typeString, "0") {
		b.WriteString(num)
	}
	if strings.Contains(typeString, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(typeString, "A") {
		b.WriteString(upper)
	}
	var tempStr = b.String()
	if len(b.String()) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < length; i++ {
		b.WriteByte(tempStr[rand.Intn(len(tempStr))])
	}
	return b.String()
}
