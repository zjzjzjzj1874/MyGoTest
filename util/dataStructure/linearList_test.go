package dataStructure

import (
	"fmt"
	"testing"
)

func TestInitLinearList(t *testing.T) {
	l := InitLinearList(5)
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
	ll := InitLinearList(5)
	// 获取元素
	fmt.Println(ll.GetElem(0))
	fmt.Println(ll.GetElem(1))
	fmt.Println(ll.GetElem(5))
	fmt.Println(ll.GetElem(10))
}

func TestLinearList_Append2(t *testing.T) {
	l := InitLinearList(5)
	fmt.Println(l)
	fmt.Println(l.Append(1))
	fmt.Println(l)
	fmt.Println(l.Append(2))
	fmt.Println(l)
	fmt.Println(l.Append(3))
	fmt.Println(l)
	fmt.Println(l.Append(4))
	fmt.Println(l)
	fmt.Println(l.Append(5))
	fmt.Println(l)
	fmt.Println(l.Append(6))
	fmt.Println(l)
}

func TestLinearList_Insert(t *testing.T) {
	l := InitLinearList(10)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	//l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Insert(4,-10))
	fmt.Println(l.Insert(7,0))
	fmt.Println(l)
}

func TestLinearList_DelLast(t *testing.T) {
	l := InitLinearList(10)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	//l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.DelLast())
	fmt.Println(l)
	fmt.Println(l.DelLast())
	fmt.Println(l)
}
func TestLinearList_Delete(t *testing.T) {
	l := InitLinearList(10)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Delete(4))
	fmt.Println(l.Delete(4))
	fmt.Println(l.Delete(4))
	fmt.Println(l.Delete(4))
	fmt.Println(l)
}
// 测试修改元素
func TestLinearList_Update(t *testing.T) {
	l := InitLinearList(10)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	//l.Append(7)
	//l.Append(8)
	//l.Append(9)
	//l.Append(10)
	fmt.Println(l)
	fmt.Println(l.Update(4,44))
	fmt.Println(l)
}

func TestLinearList_Extend(t *testing.T) {
	l := InitLinearList(5)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l)
	fmt.Println(l.Extend(0))		// expert maxsize = 10
	fmt.Println(l.Extend(5))		// expert maxsize = 15
}

func TestLinearList_Combine(t *testing.T) {
	l := InitLinearList(3)
	l.Append(1)
	l.Append(3)
	l.Append(5)

	l2 := InitLinearList(3)
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)

	fmt.Println("l1 ==== ",l)
	fmt.Println("l2 ==== ",l2)
	l3 := l.Combine(l2)
	fmt.Println(l3)
	fmt.Println(l3.Combine(l2))
}