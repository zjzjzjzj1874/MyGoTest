package interview

import (
	"fmt"
)

type Sep interface {
	Q1() // 第一题
	Q2() // 第二题
	Q3() // 第三题
}

type Day29 struct{} // 29号

func (d *Day29) Q1() {
	var ch1 chan int         // 申明channel
	ch2 := make(chan int)    // 申明channel -- 无缓冲
	ch3 := make(chan int, 3) // 申明channel -- 带缓冲
	ch4 := new(chan int)     // 申明channel

	//<- ch2	// 读取chan
	//ch1 <- 1	// 写入channel

	fmt.Println(ch1, ch2, ch3, ch4)
	fmt.Println(&ch1, &ch2, &ch3)
	fmt.Println(*ch4)
}
func (d *Day29) Q2() {
	type person struct {
		Name string
	}
	m := make(map[person]int)
	var p = person{"Jack"}
	var p1 = person{"John"}
	fmt.Println(m[p], m[p1]) // 输出值为=> 0,0
}

func hello(num ...int) {
	num[0] = 18
}
func (d *Day29) Q3() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}
