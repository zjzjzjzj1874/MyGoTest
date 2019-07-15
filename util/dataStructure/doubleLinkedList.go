package dataStructure

import (
	"fmt"
	"sort"
	"sync"
)

type (
	doubleLinkedList struct {
		// 小写的目的:不允许包外初始化:只提供New的方法初始化
		Prev *doubleLinkedList // 指针域:直接前趋地址
		Data Elem              // 数据域
		Next *doubleLinkedList // 指针域:直接后继地址
		mutex *sync.RWMutex
	}
)

// 初始化
func NewDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{
		mutex:new(sync.RWMutex),
		Prev: nil,
		Data: zero,
		Next: nil,
	}
}
// 判断空链表
func (d *doubleLinkedList) IsEmpty() bool {
	return d.Next == nil && d.Prev == nil
}
// 双向链表长度
func (d *doubleLinkedList) Len() int {
	if d.IsEmpty() {
		return 0
	}
	var length int
	node := d.Next
	for node != nil {
		length++
		node = node.Next
	}
	return length
}
// 是否包含
func (d *doubleLinkedList) Contains(elem Elem) bool {
	temp := d.Next
	for temp != nil{
		if temp.Data == elem{
			return true
		}
		temp = temp.Next
	}
	return false
}
// 转化为数组
func (d *doubleLinkedList) ToArray() (arr []int) {
	if d.IsEmpty() {
		return
	}
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	node := d.Next
	for node != nil{
		arr = append(arr, int(node.Data))
		node = node.Next
	}
	return
}
// 转化为有序数组
func (d *doubleLinkedList) SortArray() (arr []int) {
	if d.IsEmpty() {
		return
	}
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	node := d.Next
	for node != nil{
		arr = append(arr, int(node.Data))
		node = node.Next
	}
	sort.Ints(arr)
	return
}

// 增加元素
func (d *doubleLinkedList) AddFirst(elem Elem) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	temp := d.Next // 获取第一个节点
	n := doubleLinkedList{Data: elem, Prev: d, Next: temp}
	if temp != nil{
		temp.Prev = d
	}
	d.Next = &n
}
func (d *doubleLinkedList) AddLast(elem Elem)  {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	temp := d
	for temp.Next != nil{
		temp = temp.Next
	}
	n := doubleLinkedList{
		Prev:temp,
		Data:elem,
		Next:nil,
	}
	temp.Next = &n
}

// 删除元素
func (d *doubleLinkedList) RemoveFirst() bool {
	if d.IsEmpty(){
		return false
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.Next = d.Next.Next
	if d.Next != nil{
		d.Next.Prev = d
	}
	return true
}
func (d *doubleLinkedList) RemoveLast() bool {
	if d.IsEmpty(){
		return false
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	temp := d
	for temp.Next.Next != nil{
		temp = temp.Next
	}
	temp.Next = nil
	return true
}


// 遍历打印双向链表
func (d *doubleLinkedList) Traverse() {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	node := d.Next
	for node != nil {
		fmt.Printf("data is %v\n", node.Data)
		node = node.Next
	}
}
