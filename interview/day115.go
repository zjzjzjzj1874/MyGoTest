package main

import "fmt"

var y int

func f(x int) int {
	return 7
}

func day115() {
	// switch 和case的用法

	// A ==> false : " y = f(2) " used as value
	//switch y = f(2) {
	//case y == 7:
	//	return
	//}

	// B ==> ok
	//switch y = f(2); {
	//case y == 7:
	//	return
	//}

	// C ==> false : 去掉 " y = "即可
	//switch y = f(2) {
	//switch f(2) {
	//case 7:
	//	return
	//}

	// D ==> false 参考B
	//switch y = f(2); {
	//case 7:
	//	return
	//}

	// 不同切片的指针
	//a := []int{1, 2, 3, 4}
	//b := variadic(a...)
	//b[0], b[1] = b[1], b[0]
	//fmt.Println(&a,&b)
	//fmt.Println(a)

	//c := new([]int)
	//d := new([]int)
	//fmt.Println(c,d)
	//*c = append(*c, 1)
	//*c = append(*c, 2)
	//*c = append(*c, 3)
	//*c = append(*c, 4)
	//*d = variadic(*c...)
	//(*d)[0],(*d)[1] = (*d)[1],(*d)[0]
	//fmt.Println(*c)

	// 不同切片的指针
	//a := []int{1, 2, 3, 4}
	//b := NoSugar(a)
	//fmt.Println(a, b)
	//b[0], b[1] = b[1], b[0]
	//fmt.Println(a, b)

	// 切片的更多用法 ==> 可变参数传入slice的时候,
	// 函数中修改会影响到原slice(不发生扩容部分);扩容部分不会改变原始值 !!!Attention!!!
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)

	//slice := []string{"Hello","World","Go"}
	//change1(slice...)
	//fmt.Println(slice)

}

// 这种是语法糖,不会新建一个切片
func variadic(ints ...int) []int {
	return ints
}

func NoSugar(ints []int) []int {
	return ints
}

func change1(s ...string) {
	s[0] = "seekload.net"
	s = append(s, " gogogo")
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}

// 这个关于切片的案例非常经典
func change(news ...int) {
	news = append(news, 3)
	news = append(news, 4)
	news = append(news, 5)
	news = append(news, 6)
	//news[5] = 100
	//news[4] = 99 // 这里为啥不改变s的值啊???
	news[4] = 98 // 这里为啥不改变外面s的值啊???
	fmt.Println(news)
}
