package dataStructure

import "fmt"

type (
	doubleLinkedList struct{
		Prev 	*doubleLinkedList	// 指针域:直接前趋地址
		Data	Elem				// 数据域
		Next	*doubleLinkedList	// 指针域:直接后继地址
	}
)

// 初始化
func NewDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{
		Prev: nil,
		Data: zero,
		Next: nil,
	}
}

func (d *doubleLinkedList) Append(elem Elem) {
	n := doubleLinkedList{Data:elem,Prev:d,Next:nil}
	d.Next = &n
}

func (d *doubleLinkedList) Traverse() {
	ptr := d.Next
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}
