package util

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println(Get("www.baidu.com"))
	//fmt.Println(Get("http://www.baidu.com"))
}