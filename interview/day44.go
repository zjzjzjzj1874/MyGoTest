package main

import (
	"fmt"
	"reflect"
)

func day44() {

	//lenAndCap()

	//nilTest()

	typeVar()

}

// region 数组与切片的属性
func lenAndCap() {
	// -------------------------------------------------
	// -------------------- 数组 -----------------------
	// -------------------------------------------------
	arr1 := [5]int{1, 2, 3}   // 指定数组的长度
	arr2 := [...]int{1, 2, 3} // 自动推断数组长度
	fmt.Println(arr1, "\n", arr2)

	// 按索引给未指定长度数组赋值 ==> 中间或前面的位置补0
	arr3 := [...]int{2: 4, 4: 4, 3}
	fmt.Println("按索引给数组赋值1=>", arr3)
	// 按索引给定长数组赋值 ==> 不可超过定长
	arr4 := [5]int{2: 2, 4: 4}
	//arr4 := [5]int{2: 2, 4: 4, 5} // ==> 会报错:因为已经超过数组的长度了
	fmt.Println("按索引给数组赋值2=>", arr4)

	// -------------------------------------------------
	// -------------------- 切片 -----------------------
	// -------------------------------------------------
	// 初始化切片
	var s1 []int
	s2 := make([]int, 5, 5+3)
	fmt.Println("切片初始化:s1 == ", s1, "\n s2 == ", s2)
	fmt.Printf("s1 长度:%d,容量:%d \n", len(s1), cap(s1))
	fmt.Printf("s2 长度:%d,容量:%d \n", len(s2), cap(s2))

	// -------------------------------------------------
	// -------------------- 总结 -----------------------
	// -------------------------------------------------
	// 1.数组和切片定义不一样;
	// 2.数组和切片初始化,申明不一样;
	// 3.数组是定长的和切片初始化,申明不一样;
	// 4.数组是值传递,切片是地址传递;==>函数中的值拷贝
	// 5.注意,数组不能使用append的方法=>是切片的内置方法

	// 1.下面代码有何问题
	//m := make(map[string]int, 2) // 注意:使用make创建map,第二个参数会被忽略
	//cap(m)
	// err :不能使用cap()方法获取map的容量或长度;cap()适用于数组,数组指针,slice和channel

	// 其实 Go 语言的赋值和函数传参规则很简单，除了闭包函数以引用的方式对外部变量访问之外，其它赋值和函数传参数都是以传值的方式处理。
}

// endregion 数组与切片的属性

// region nil的零值与初始化
func nilTest() {
	var x interface{} = nil // 正确的类型
	// var x = nil // 错误=>not explicit type==>未指明类型;静态强类型语言.
	// nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。
	// 如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误。
	fmt.Println("x的类型是 === ", reflect.TypeOf(x))
	_ = x
}

// endregion nil的零值与初始化

// region 变量申明测试
type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}
func typeVar() {
	var data info
	var err error
	// data.result,err := work() // 错误=>不能使用短变量申明来设置结构体字段值 ==> Can't Use Short Variable Declarations to Set Field Values
	data.result, err = work()
	fmt.Printf("result:%+v;err=> %+v \n", data.result, err)
}

// endregion 变量申明测试
