package util

import (
	"flag"
	"fmt"
)

type Flag struct {
	Name     string // flag在命令行中的名字
	Usage    string // 帮助信息
	Value    string // 要设置的值
	DefValue string // 默认值（文本格式），用于使用信息
}

func MyFlag() {
	var ip = flag.String("flagName", "flagValue", "help message for flagName")
	flag.Parse()

	fmt.Println("ip has value ==> ", *ip)

	var gopherType string

	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+"(shorthand)")
	fmt.Println(flag.Args())
	fmt.Println(gopherType)

}
