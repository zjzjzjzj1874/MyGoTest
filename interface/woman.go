package people

import "fmt"

// Woman结构体,实现了People接口中所有的方法, ==> Woman就属于People这个类型,这个接口了
type Woman struct {
	Name string
	Age  int
}

func (w *Woman) Say() {
	fmt.Println("hello,I'w a woman,my name is", w.Name)
}
func (w *Woman) Eat() {
	fmt.Println("I like eating apple")
}
func (w *Woman) Sleep() {
	fmt.Println("Sleeping makes me beautiful")
}
