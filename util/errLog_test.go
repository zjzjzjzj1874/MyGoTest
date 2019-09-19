package util

import (
	"fmt"
	"testing"
)

func TestErrorLog(t *testing.T) {
	phone := "15822334455"
	fmt.Println(phone[len(phone)-4:])
}

func TestInfoLog(t *testing.T) {
	//InfoLog("我","hello world")
	fmt.Println(Config.LogSuffix)

	a := make(map[int]interface{})
	a[3] = [2]int{}
	a[3] = 45
}
