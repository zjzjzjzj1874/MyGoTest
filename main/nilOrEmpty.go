package main

import "fmt"

type I interface {
	m()
}
type T struct {
	num int
}

//T类型的一个方法,空方法
func (t T) m() {
	fmt.Println("I’m alive!")
}

type P *T

func main() {
	/*var p = &T{num: 10}
	fmt.Println(p.num)
	(*p).m()
	p.m()*/

	//str := "我是"
	//fmt.Printf("\r\n %v 123",str)
	//fmt.Println("\r\n 123",str)

	TestNilOrEmpty()
}

//测试interface的nil和空
func Combine() {
	var i I //1:申明 type为I的变量:i
	i = T{} //2:赋值	i = 空的T {0}

	//上面两步可以一步执行;
	i1 := I(T{})
	i1.m()

	i.m()
}

//测试nil和empty的方法
func TestNilOrEmpty() {
	p := T{}           //这样申明:表示 p = T{0}  是一个空的T类型
	fmt.Println(p)     //打印值:{0}
	fmt.Println(p.num) //打印值:0
	p.m()              //没有问题,不会报错

	var p1 *T //这样申明:是一个nil的T,因为没有初始化
	p1 = &p
	fmt.Println(p1, ";address p1 == ", &p1)
	fmt.Println(p, ";address   p == ", &p)
	fmt.Println(p, ";p == ", *p1)
	p.m() //没问题,不报错
	//fmt.Println(p1.num)	//报错:p1是nil的,不存在num的属性;   ---- panic: runtime error: invalid memory address or nil pointer dereference

}
