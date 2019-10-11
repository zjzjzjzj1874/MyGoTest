package util

import (
	"flag"
	"fmt"
	"testing"
)

func TestMyFlag(t *testing.T) {
	// 注册参数
	port := flag.Int("p", 8080, "server Port")
	// 解析参数, 模块方法将使用 os.Args[1:] 做参数解析
	flag.Parse()
	fmt.Printf("server port:%d", *port)
	fmt.Printf("server port address:%d", port)
}
