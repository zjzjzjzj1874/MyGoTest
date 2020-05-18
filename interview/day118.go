package main

import "fmt"

func day118() {
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
	// 在changes里面,当append把s撑到扩容后,原来的s切片里面是12345,
	// 但是扩容后的切片会随着操作而改变其值;这个时候是默认操作扩容后的切片,
	// 原切片将和结构体一样,只是简单的浅拷贝.原始值不会改变了
	s[4] = 98 // 这里为啥不改变外面s的值啊???
	fmt.Println(s)
}
