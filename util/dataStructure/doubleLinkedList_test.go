package dataStructure

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList_Append(t *testing.T) {
	d := NewDoubleLinkedList()
	fmt.Println(d.ToArray())
	fmt.Println(d.IsEmpty())
	fmt.Println(d.Len())
	d.AddFirst(1)
	fmt.Println(d.IsEmpty())
	fmt.Println(d.Len())
	d.AddFirst(3)
	fmt.Println(d.Len())
	fmt.Println(d.IsEmpty())
}

func TestDoubleLinkedList_AddLast(t *testing.T) {
	d := NewDoubleLinkedList()
	d.AddLast(2)
	d.AddLast(4)
	d.AddLast(6)
	d.Traverse()
}
func TestDoubleLinkedList_Contains(t *testing.T) {
	d := NewDoubleLinkedList()
	d.AddLast(2)
	d.AddLast(4)
	d.AddLast(6)
	fmt.Println(d.Contains(2))
	fmt.Println(d.Contains(3))
}

func TestDoubleLinkedList_ToArray(t *testing.T) {
	d := NewDoubleLinkedList()
	d.AddFirst(1)
	d.AddFirst(3)
	d.AddFirst(5)
	d.AddFirst(7)
	fmt.Println(d.ToArray())
	fmt.Println(d.SortArray())
}
func TestDoubleLinkedList_RemoveFirst(t *testing.T) {
	d := NewDoubleLinkedList()
	fmt.Println(d.RemoveFirst())
	d.AddLast(2)
	d.AddLast(4)
	d.RemoveFirst()
	fmt.Println(d.ToArray())
}
func TestDoubleLinkedList_RemoveLast(t *testing.T) {
	d := NewDoubleLinkedList()
	d.AddLast(2)
	d.AddLast(4)
	d.AddLast(6)
	fmt.Println(d.RemoveLast())
	fmt.Println(d.RemoveLast())
	//fmt.Println(d.RemoveLast())
	//fmt.Println(d.RemoveLast())
	fmt.Println(d.ToArray())
}