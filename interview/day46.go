package main

import "fmt"

func main() {
	//constant()
	stringNil()
}

// region 常量申明
func constant() {
	// region 例1 true
	const x = 123
	const y = 1.23
	fmt.Println(x)
	// endregion 例1 true

	// region 例2 true
	const (
		a uint16 = 120
		b
		c = "abc"
		d
	)
	fmt.Printf("%T %v\n", b, b) // uint16 120
	fmt.Printf("%T %v\n", d, d) // string abc
	// endregion 例2 true
}

// endregion 常量申明

// region string的默认值
func stringNil() {
	var x string = ""
	//var x string = nil // 错误 => 初始化申明失败=> string的默认值是 ""
	if x == "" {
		x = "default"
	}
	fmt.Println(x)
}

// endregion string的默认值
