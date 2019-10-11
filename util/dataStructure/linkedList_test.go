package dataStructure

import (
	"container/list"
	"fmt"
	"testing"
)

func initLinkedList() *linkedList {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(3)
	l.AddLast(5)
	l.AddLast(2)
	l.AddLast(4)
	l.AddLast(6)
	return l
}

// 验证是否为空
func TestLinkedList_IsEmpty(t *testing.T) {
	l := NewLinkedList()
	fmt.Println(l.IsEmpty())
	l.Traverse()
}

func TestSequenceList_AddLast(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(3)
	l.AddLast(5)
	l.Traverse()
}

func TestLinkedList_AddFirst(t *testing.T) {
	l := NewLinkedList()
	l.AddFirst(1)
	l.AddFirst(3)
	l.AddFirst(5)
	l.Traverse()
}

func TestLinkedList_Add(t *testing.T) {
	l := NewLinkedList()
	l.AddLast(1)
	l.AddLast(3)
	l.AddLast(5)
	l.Traverse()
	l.Add(0, 4)
	fmt.Println("after ===== ")
	l.Traverse()
}

func TestLinkedList_Remove(t *testing.T) {
	l := initLinkedList()
	l.Traverse()
	fmt.Println("链表长度 == ", l.Len())
	fmt.Println(l.Remove(3))
	fmt.Println("after ===== ")
	l.Traverse()
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	l := initLinkedList()
	fmt.Println(l.RemoveFirst())
	fmt.Println(l.RemoveFirst())
	l.Traverse()
}

func TestLinkedList_RemoveLast(t *testing.T) {
	l := initLinkedList()
	fmt.Println(l.RemoveLast())
	fmt.Println(l.RemoveLast())
	l.Traverse()
}

func TestLinkedList_ToArray(t *testing.T) {
	l := initLinkedList()
	//fmt.Println(l.ToArray())
	fmt.Println(l.ToSortArray())
}

func TestLinkedList_Contains(t *testing.T) {
	l := initLinkedList()
	fmt.Println(l.Contains(1))
	fmt.Println(l.Contains(4))
	fmt.Println(l.Contains(6))
	fmt.Println(l.Contains(7))
	fmt.Println(l.ToSortArray())
	fmt.Println(l.ToArray())
}

// 修改元素
func TestLinkedList_SetFirst(t *testing.T) {
	l := initLinkedList()
	l.Clear()
	fmt.Println("清空之后的链表 === ", l.ToArray())
	l.SetFirst(100)
	fmt.Println(l.ToArray())
}

func TestLinkedList_SetLast(t *testing.T) {
	l := initLinkedList()
	fmt.Println(l.ToArray())
	l.SetLast(100)
	fmt.Println(l.ToArray())
}
func TestLinkedList_Set(t *testing.T) {
	l := initLinkedList()
	fmt.Println(l.ToArray())
	fmt.Println(l.Set(7, 3))
	fmt.Println(l.Set(3, 3))
	fmt.Println(l.ToArray())
}

func TestLinkedList_GetFirst(t *testing.T) {
	l := initLinkedList()
	//l.Clear()
	fmt.Println(l.ToArray())
	fmt.Println(l.GetFirst())
}

func TestLinkedList_GetLast(t *testing.T) {
	l := initLinkedList()
	//l.Clear()
	fmt.Println(l.ToArray())
	fmt.Println(l.GetLast())
}
func TestLinkedList_Get(t *testing.T) {
	l := initLinkedList()
	//l.Clear()
	fmt.Println(l.ToArray())
	fmt.Println(l.Get(0))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(5))
	fmt.Println(l.Get(7))
}

// go语言自带list实现
func TestGOList(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
