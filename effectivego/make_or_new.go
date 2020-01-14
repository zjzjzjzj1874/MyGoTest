package main

import "fmt"

// 测试new和var的作用与区别
func newOrVar() {
	v1 := new([]int)
	var v2 []int
	v3 := make([]int, 0)
	fmt.Println(*v1, v2, &v3)

	*v1 = append(*v1, 1)

	v2 = append(v2, 2)

	v3 = append(v3, 3)

	fmt.Println(v1, "\n", v2, "\n", v3)
}

func main() {
	newOrVar()
}
