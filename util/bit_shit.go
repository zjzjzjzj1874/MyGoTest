package util

import "fmt"

// 位移动
func LeftShift() {
	const (
		a = 1 << iota // 1 << 0
		b             // 1 << 1
		c             // 1 << 2
		d             // 1 << 3
		e = 1 << 6
	)
	fmt.Println(a) // 1
	fmt.Println(b) // 2
	fmt.Println(c) // 4
	fmt.Println(d) // 8
	fmt.Println(e) // 64

	// 左移位运算可以看做乘法
	fmt.Println(30 << 3)
	fmt.Println(2 << 3)
	fmt.Println(5 << 3)
}

func RightShift() {
	const (
		a = 16 >> iota // 16 >> 0
		b              // 16 >> 1
		c              // 16 >> 2
		d              // 16 >> 3
	)
	fmt.Println(a) // 16
	fmt.Println(b) // 8
	fmt.Println(c) // 4
	fmt.Println(d) // 2

	// 右移位运算可以看做除法
	fmt.Println(32 >> 3)
	fmt.Println(8 >> 3)
	fmt.Println(5 >> 3)
}
