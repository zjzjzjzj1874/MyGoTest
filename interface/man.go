package people

import "fmt"

// Man结构体,实现了People接口中所有的方法, ==> Man就属于People这个类型,这个接口了
type Man struct {
	Name string
	Age  int
}

func (m *Man) Say() {
	fmt.Println("hello,I'm a man,my name is", m.Name)
}
func (m *Man) Eat() {
	fmt.Println("I like eating banana")
}
func (m *Man) Sleep() {
	fmt.Println("I'm going to sleep")
}
