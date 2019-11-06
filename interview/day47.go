package main

import "fmt"

func main() {

	//preBuild()

	openingBrace()

}

// region 内建前置
func preBuild() {
	data := []int{1, 2, 3}
	i := 0
	i++
	//++i // syntax error: unexpected ++, expecting }
	//fmt.Println(data[i++]) // syntax error: unexpected ++, expecting :
	fmt.Println(data[i]) // ok
}

// endregion 内建前置

// region 变量申明与左大括号
func openingBrace() {
	x := 1
	fmt.Println("first x == ", x)
	// 在函数,if,else等关键词后面的左大括号不可换行
	{ // 注意,这里左大括号是不会报错的 ==> 这样子我可以控制变量的生命周期了
		fmt.Println("second x == ", x)
		i, x := 2, 2
		fmt.Println("third x == ", i, x)
	}
	fmt.Println("forth x == ", x)
}

// endregion
