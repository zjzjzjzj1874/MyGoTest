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
	LinearElem int

	// 定义一个线性表结构体
	LinearList struct {
		Lock    sync.Mutex   // 互斥锁 使用方法: 加锁 = l.Lock.Lock()和解锁 = l.Lock().Unlock()
		Maxsize int          // 最大长度
		CurLen  int          // 当前长度
		Data    []LinearElem // 保存数据(也可以把linearElem换成interface{},更加通用)
	}
)

func (l *LinearList) Len() int {
	return l.CurLen
}

func (l *LinearList) Less(i, j int) bool {
	return l.Data[i] < l.Data[j]
}

func (l *LinearList) Swap(i, j int) {
	l.Data[i], l.Data[j] = l.Data[j], l.Data[i]
}

// 初始化一个线性表
func NewList(maxsize int) *LinearList {
	if maxsize < one {
		maxsize = 10
	}
	return &LinearList{Maxsize: maxsize, Data: make([]LinearElem, maxsize)}
}

// 判断线性表是否为空
func (l *LinearList) IsEmpty() bool {
	return zero == l.CurLen
}

// 判断线性表是否已满
func (l *LinearList) IsFull() bool {
	return len(l.Data) == l.CurLen
}

// 判断元素是否在线性表里
func (l *LinearList) Contains(elem LinearElem) bool {
	for i := zero; i < l.CurLen; i++ {
		if elem == l.Data[i] {
			return true
		}
	}
	return false
}

// 判断索引是否越界
func (l *LinearList) IndexOutOfRange(index int) bool {
	//return index > l.Maxsize || index < 1
	return index > len(l.Data) || index < one
}

// 根据index获取对应元素
func (l *LinearList) Get(index int) (LinearElem, bool) {
	if l.IndexOutOfRange(index) {
		return -one, false
	}
	return l.Data[index-one], true
}

// 插入元素
// 1.末尾append一个数据 当前长度++
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) AddLast(ele LinearElem) int {
	if l.IsFull() {
		return -one
	}
	l.Data[l.CurLen] = ele
	l.CurLen++
	return l.CurLen
}

// 2.在index位置前插入元素e;
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) Add(index int, e LinearElem) int {
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
func (l *LinearList) RemoveLast() int {
	if l.IsEmpty() {
		return -one
	}
	l.Data[l.CurLen-one] = 0
	l.CurLen--
	return l.CurLen
}

// 2.删除指定位置元素
// 返回值:-1失败  成功后返回:删除后的线性表长度
func (l *LinearList) Remove(index int) int {
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
func (l *LinearList) Set(index int, ele LinearElem) int {
	if index > l.CurLen || index < one {
		return -one
	}
	l.Data[index-one] = ele
	return index
}

// 扩容:i<1的时候按规则扩容,指定长度的话,按照指定长度扩容
// 扩容规则:总长度小于100的时候,扩容两倍;其他扩容1.5倍
func (l *LinearList) Extend(i int) *LinearList {
	maxsize := len(l.Data)
	switch {
	case i < one && maxsize < doubleTippingPoint:
		maxsize *= double
	case i < one && maxsize >= doubleTippingPoint:
		maxsize += (maxsize%two + maxsize) / two
	default:
		maxsize += i
	}
	extend := make([]LinearElem, maxsize-len(l.Data))
	l.Data = append(l.Data, extend...)
	l.Maxsize = len(l.Data)
	return l
}

// 两个线性表的合并
func (l *LinearList) Combine(l2 *LinearList) *LinearList {
	l2Len := make([]LinearElem, len(l2.Data))
	l.Data = append(l.Data, l2Len...)
	l.Maxsize = len(l.Data) + len(l2.Data)
	for i := zero; i < l2.CurLen; i++ {
		l.AddLast(l2.Data[i])
	}
	return l
}

// 线性表排序
// 从小到大排序 -- 从大到小排序:就修改Less方法
func (l *LinearList) SortList() *LinearList {
	sort.Stable(l)
	return l
}

// 清空线性表
func (l *LinearList) Clear() *LinearList {
	for i := zero; i < l.CurLen; i++ {
		l.Data[i] = zero
	}
	l.CurLen = zero
	return l
}
