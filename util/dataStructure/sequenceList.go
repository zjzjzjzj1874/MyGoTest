package dataStructure

import (
	"sort"
	"sync"
)

// TODO 待完成:
//  1.线性表中存储数据不仅限于int
// TODO 待完善方法:addFirst,addLast,
//  getFirst,getLast,
//  removeFirst,removeLast,removeRange(删除某个范围内元素),
//  toArray(转化为数组,数组长度为线性表长度),
//  查找元素位置:indexOf,lastIndexOf

const (
	doubleTippingPoint = 100
	double             = 2
)

const (
	zero = iota
	one
	two
)

type (
	// 线性表中的元素
	SequenceElem int

	// 定义一个线性表结构体
	SequenceList struct {
		Lock    sync.Mutex   // 互斥锁 使用方法: 加锁 = l.Lock.Lock()和解锁 = l.Lock().Unlock()
		Maxsize int          // 最大长度
		CurLen  int          // 当前长度
		Data    []SequenceElem // 保存数据(也可以把linearElem换成interface{},更加通用)
	}
)

func (l *SequenceList) Len() int {
	return l.CurLen
}

func (l *SequenceList) Less(i, j int) bool {
	return l.Data[i] < l.Data[j]
}

func (l *SequenceList) Swap(i, j int) {
	l.Data[i], l.Data[j] = l.Data[j], l.Data[i]
}

// 初始化一个线性表
func NewList(maxsize int) *SequenceList {
	if maxsize < one {
		maxsize = 10
	}
	return &SequenceList{Maxsize: maxsize, Data: make([]SequenceElem, maxsize)}
}

// 线性表转化为数组
func (l *SequenceList) ToArray() []int {
	var arr []int
	for i:=0;i<l.CurLen;i++{
		arr = append(arr, int(l.Data[i]))
	}
	return arr
}
// 线性表转化为有序数组
func (l *SequenceList) ToSortArray() []int {
	var arr []int
	for i:=0;i<l.CurLen;i++{
		arr = append(arr, int(l.Data[i]))
	}
	sort.Ints(arr)
	return arr
}


// 判断线性表是否为空
func (l *SequenceList) IsEmpty() bool {
	return zero == l.CurLen
}

// 判断线性表是否已满
func (l *SequenceList) IsFull() bool {
	return len(l.Data) == l.CurLen
}

// 判断元素是否在线性表里
func (l *SequenceList) Contains(elem SequenceElem) bool {
	for i := zero; i < l.CurLen; i++ {
		if elem == l.Data[i] {
			return true
		}
	}
	return false
}

// 判断索引是否越界
func (l *SequenceList) IndexOutOfRange(index int) bool {
	//return index > l.Maxsize || index < 1
	return index > len(l.Data) || index < one
}

// 根据index获取对应元素
func (l *SequenceList) Get(index int) (SequenceElem, bool) {
	if l.IndexOutOfRange(index) {
		return -one, false
	}
	return l.Data[index-one], true
}

// 插入元素
// 1.末尾append一个数据 当前长度++
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *SequenceList) AddLast(ele SequenceElem) int {
	if l.IsFull() {
		return -one
	}
	l.Data[l.CurLen] = ele
	l.CurLen++
	return l.CurLen
}

// 2.在index位置前插入元素e;
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *SequenceList) Add(index int, e SequenceElem) int {
	if l.IsFull() || index < one || index > l.CurLen {
		return -one
	}
	for i := l.CurLen; i > index-one; i-- {
		l.Data[i] = l.Data[i-one]
	}
	l.Data[index-one] = e
	l.CurLen++
	return l.CurLen
}

// 删除元素
// 1.删除最后一个元素
// 返回值 -1失败;成功后返回:删除后的线性表长度
func (l *SequenceList) RemoveLast() int {
	if l.IsEmpty() {
		return -one
	}
	l.Data[l.CurLen-one] = 0
	l.CurLen--
	return l.CurLen
}

// 2.删除指定位置元素
// 返回值:-1失败  成功后返回:删除后的线性表长度
func (l *SequenceList) Remove(index int) int {
	if l.IsEmpty() || index > l.CurLen || index < one {
		return -one
	}
	for i := index; i < l.CurLen; i++ {
		l.Data[i-one] = l.Data[i]
	}
	l.Data[l.CurLen-one] = 0
	l.CurLen--
	return l.CurLen
}

// 修改替换
func (l *SequenceList) Set(index int, ele SequenceElem) int {
	if index > l.CurLen || index < one {
		return -one
	}
	l.Data[index-one] = ele
	return index
}

// 扩容:i<1的时候按规则扩容,指定长度的话,按照指定长度扩容
// 扩容规则:总长度小于100的时候,扩容两倍;其他扩容1.5倍
func (l *SequenceList) Extend(i int) *SequenceList {
	maxsize := len(l.Data)
	switch {
	case i < one && maxsize < doubleTippingPoint:
		maxsize *= double
	case i < one && maxsize >= doubleTippingPoint:
		maxsize += (maxsize%two + maxsize) / two
	default:
		maxsize += i
	}
	extend := make([]SequenceElem, maxsize-len(l.Data))
	l.Data = append(l.Data, extend...)
	l.Maxsize = len(l.Data)
	return l
}

// 两个线性表的合并
func (l *SequenceList) Combine(l2 *SequenceList) *SequenceList {
	l2Len := make([]SequenceElem, len(l2.Data))
	l.Data = append(l.Data, l2Len...)
	l.Maxsize = len(l.Data) + len(l2.Data)
	for i := zero; i < l2.CurLen; i++ {
		l.AddLast(l2.Data[i])
	}
	return l
}

// 线性表排序
// 从小到大排序 -- 从大到小排序:就修改Less方法
func (l *SequenceList) SortList() *SequenceList {
	sort.Stable(l)
	return l
}

// 清空线性表
func (l *SequenceList) Clear() *SequenceList {
	for i := zero; i < l.CurLen; i++ {
		l.Data[i] = zero
	}
	l.CurLen = zero
	return l
}
