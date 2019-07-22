package log

import "testing"

func TestErrorLog(t *testing.T) {
	ErrorLog("测试添加","123","456")
}

func TestInfoLog(t *testing.T) {
	InfoLog("我","hello world")
}
