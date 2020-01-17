package main

import "fmt"

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	changes(slice...)
	fmt.Println(slice)
	changes(slice[0:2]...)
	fmt.Println(slice)
}

// 这个关于切片的案例非常经典
func changes(s ...int) {
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)
	s = append(s, 6)
	//s[5] = 100
	//s[4] = 99 // 这里为啥不改变s的值啊???
	s[4] = 98 // 这里为啥不改变外面s的值啊???
	fmt.Println(s)
}
