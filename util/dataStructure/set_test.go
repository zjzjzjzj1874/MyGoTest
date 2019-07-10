package dataStructure

import (
	"fmt"
	"testing"
)

func newSet2() *Set {
	s := NewSet()
	s.Add(2)
	s.Add(4)
	s.Add(6)
	s.Add(8)
	s.Add(10)
	return s
}

func TestNewSet(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(4)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	fmt.Println(s.SortList())
	//
	//	// remove测试
	//	//s.Remove(5)
	//	//fmt.Println(s)
	//	//s.Remove(6)
	//	//fmt.Println(s)
	//
	//	// contains测试
	//	//fmt.Println(s.Contains(1))
	//	//fmt.Println(s.Contains(6))
	//
	//	// clear清空集合
	//	//s.Clear()
	//fmt.Println(s)

	// 转化为list,sortList
	//fmt.Println(s.List())
	//fmt.Println(s.SortList())
	//fmt.Println(s.Len())

	// 集合的运算
	s2 := newSet2()
	fmt.Println(s2.SortList())
	//fmt.Println(s.Union(s2).SortList())	// 并集
	//fmt.Println(s.InterSection(s2).SortList())	// 交集
	fmt.Println(s.Equal(s2)) // 判断两个集合是否相等
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet()
	fmt.Println(s.IsEmpty())
	s2 := newSet2()
	fmt.Println(s2.IsEmpty())
}

func TestSet_IsSubSet(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(5)
	s2 := NewSet()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	fmt.Println(s.List())
	fmt.Println(s.SortList())
	fmt.Println(s.IsSubSet(s2))
}
// 差集测试
func TestSet_DiffSet(t *testing.T) {
	s := NewSet()
	//s.Add(1)
	//s.Add(2)
	//s.Add(3)
	s2 := NewSet()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	fmt.Println(s.SortList())
	fmt.Println(s2.SortList())
	fmt.Println(s.DiffSet(s2).SortList())
	fmt.Println(s2.DiffSet(s).SortList())
}
// 对称差集测试
func TestSet_SymDiffSet(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s2 := NewSet()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	fmt.Println(s.SortList())
	fmt.Println(s2.SortList())
	fmt.Println(s.SymDiffSet(s2).SortList())
}