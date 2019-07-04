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