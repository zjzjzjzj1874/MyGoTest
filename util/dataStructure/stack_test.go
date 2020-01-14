package dataStructure

import (
	"fmt"
	"testing"
)

func TestNewSliceStack(t *testing.T) {
	stack := NewSliceStack()

	stack.Push(12)
	stack.Push("无名之辈,我是谁")
	stack.Push("Hello,world!")

	stack.Clear()
	fmt.Println("栈顶元素:", stack.Top())
	fmt.Println("移除前:", stack)

	stack.Clear()
	stack.Pop()
	fmt.Println("移除后:", stack)

	stack.Clear()
	fmt.Println(stack.Size())
}

func TestNewLinkedStack(t *testing.T) {
	ls := NewLinkedStack()
	ls.Push(12)
	ls.Push(23)
	ls.Push(34)
	fmt.Println(ls)
	fmt.Println(ls.Elem)
}
