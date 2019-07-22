package util

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// 验证传入的str是否含有中文字符串
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

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

//将可能是int\int32\int64\float64\string类型的interface{}转换成string，可继续完善类型
func InterfaceToString(i interface{}) string {
	switch i.(type) {
	case int:
		return strconv.Itoa(i.(int))
	case int32:
		return strconv.Itoa(int(i.(int32)))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case string:
		return i.(string)
	default:
		str,_ := json.Marshal(i)
		return string(str)
	}
}