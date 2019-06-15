package util

import (
	"fmt"
	"testing"
)

// 测试指定长度的随机字符串生成
func TestGetRandomString(t *testing.T) {
	fmt.Println(GetRandomString(5, "0"))
	fmt.Println(GetRandomString(5, "a"))
	fmt.Println(GetRandomString(5, "A"))

	fmt.Println(GetRandomString(10, "0a"))
	fmt.Println(GetRandomString(10, "0A"))
	fmt.Println(GetRandomString(10, "Aa"))

	fmt.Println(GetRandomString(15, "0Aa"))
}
