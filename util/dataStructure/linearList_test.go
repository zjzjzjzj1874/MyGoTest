package dataStructure

import (
	"fmt"
	"testing"
)

func TestInitLinearList(t *testing.T) {
	l := NewList(5)
	fmt.Println(l.IsEmpty())
	fmt.Println(l.IsFull())
	fmt.Println(l.Data)
	fmt.Println(l.CurLen)
	// 判断下标是否越界
	fmt.Println(l.IndexOutOfRange(0))
	fmt.Println(l.IndexOutOfRange(2))
	fmt.Println(l.IndexOutOfRange(10))
}

func TestLinearList_GetElem(t *testing.T) {
	ll := NewList(5)
	// 获取元素
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))
	fmt.Println(ll.Get(5))
	fmt.Println(ll.Get(10))
}

func TestLinearList_Append2(t *testing.T) {
	l := NewList(5)
	fmt.Println(l)
	fmt.Println(l.AddLast(1))
	fmt.Println(l)
	fmt.Println(l.AddLast(2))
	fmt.Println(l)
	fmt.Println(l.AddLast(3))
	fmt.Println(l)
	fmt.Println(l.AddLast(4))
	fmt.Println(l)
	fmt.Println(l.AddLast(5))
	fmt.Println(l)
	fmt.Println(l.AddLast(6))
	fmt.Println(l)
}

func TestLinearList_Insert(t *testing.T) {
	l := NewList(10)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	//l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Add(4,-10))
	fmt.Println(l.Add(7,0))
	fmt.Println(l)
}

func TestLinearList_DelLast(t *testing.T) {
	l := NewList(10)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	//l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.RemoveLast())
	fmt.Println(l)
	fmt.Println(l.RemoveLast())
	fmt.Println(l)
}
func TestLinearList_Delete(t *testing.T) {
	l := NewList(10)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	l.AddLast(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Remove(4))
	fmt.Println(l.Remove(4))
	fmt.Println(l.Remove(4))
	fmt.Println(l.Remove(4))
	fmt.Println(l)
}
// 测试修改元素
func TestLinearList_Update(t *testing.T) {
	l := NewList(10)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	l.AddLast(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Set(4,44))
	fmt.Println(l)
}

func TestLinearList_Extend(t *testing.T) {
	l := NewList(5)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	fmt.Println(l)
	fmt.Println(l.Extend(0))		// expert maxsize = 10
	fmt.Println(l.Extend(5))		// expert maxsize = 15
}

func TestLinearList_Combine(t *testing.T) {
	l := NewList(3)
	l.AddLast(1)
	l.AddLast(3)
	l.AddLast(5)

	l2 := NewList(3)
	l.AddLast(2)
	l.AddLast(4)
	l.AddLast(6)

	fmt.Println("l1 ==== ",l)
	fmt.Println("l2 ==== ",l2)
	l3 := l.Combine(l2)
	fmt.Println(l3)
	fmt.Println(l3.Combine(l2))
}
// 为测试用例初始化一个线性表
func initLinearList() *LinearList {
	l := NewList(10)
	l.AddLast(2)
	l.AddLast(1)
	l.AddLast(4)
	l.AddLast(3)
	l.AddLast(6)
	l.AddLast(5)
	fmt.Println(l)
	return l
}
func TestLinearList_Has(t *testing.T) {
	l := initLinearList()
	fmt.Println(l.Contains(0))
	fmt.Println(l.Contains(1))
}

func TestLinearList_SortList(t *testing.T) {
	l := initLinearList()
	fmt.Println(l.SortList())
}

func TestLinearList_Clear(t *testing.T) {
	l := initLinearList()
	fmt.Println(l.Clear())
}