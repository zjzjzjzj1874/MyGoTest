package dataStructure

import (
	"fmt"
	"sort"
)

// TODO 是否需要查找给定元素位置的方法?
// 头结点：有时在链表的第一个结点之前会额外增设一个结点，结点的数据域一般不存放数据（有些情况下也可以存放链表的长度等信息），此结点被称为头结点。 -- 非必须
// 首元结点：链表中第一个元素所在的结点，它是头结点后边的第一个结点。
// 头指针：永远指向链表中第一个结点的位置（如果链表有头结点，头指针指向头结点；否则，头指针指向首元结点）。 -- 必须有
type (
	Elem int				// 链表数据类型
	linkedList struct {		// 单链表
		Data Elem   		// 数据域 -- 一般都是非基本数据类型
		Next *linkedList 	// 指针域
	}
)

// 初始化一个链表:链表有一个头结点,指针域为空,链表则为空表 -- 单链表可以没有头结点,但必须有头指针
func NewLinkedList() *linkedList {
	return &linkedList{Data:0,Next:nil}	// 有头结点
}

// 遍历链表 -- 打印链表
func (l *linkedList) Traverse() {
	point := l.Next
	for point != nil {
		fmt.Println(point.Data) //输出数据
		point = point.Next
	}
}

// 将链表转化为数组
func (l *linkedList) ToArray() []int {
	if l.IsEmpty(){
		return nil
	}
	var arr []int
	temp := l.Next
	for temp != nil{
		arr = append(arr, int(temp.Data))
		temp = temp.Next
	}
	return arr
}
// 将链表转化为有序数组
func (l *linkedList) ToSortArray() []int {
	if l.IsEmpty(){
		return nil
	}
	var arr []int
	temp := l.Next
	for temp != nil{
		arr = append(arr, int(temp.Data))
		temp = temp.Next
	}
	sort.Ints(arr)
	return arr
}

// 是否包含
func (l *linkedList) Contains(i Elem) bool {
	if l.IsEmpty(){
		return false
	}
	temp := l.Next
	for temp != nil{
		if temp.Data == i{
			return true
		}
		temp = temp.Next
	}
	return false
}

// 是否为空
func (l *linkedList) IsEmpty() bool {
	return l.Next == nil
}

// 清空链表
func (l *linkedList) Clear() {
	l.Next = nil
}
// 链表的长度
func (l *linkedList) Len() int {
	if l.IsEmpty(){
		return zero
	}
	point := l.Next
	var length int
	for point != nil {
		length++
		point = point.Next
	}
	return length
}

// 插入元素1:链表末尾
func (l *linkedList) AddLast(data Elem) {
	point := l
	for point.Next != nil {
		point = point.Next
	}
	var NewLinklist = linkedList{Data:data}
	point.Next = &NewLinklist
}
// 插入元素2:链表头结点后
func (l *linkedList) AddFirst(data Elem) {
	var nl = linkedList{Data:data}
	nl.Next = l.Next
	l.Next = &nl
}
// 插入元素3:指定位置
func (l *linkedList) Add(index int, data Elem) bool {
	if index < zero || index > l.Len() {
		return false
	} else {
		point := l
		for i := one; i < index; i++ {	// i=1,在index前插入;i=0,在index之后插入
			point = point.Next
		}
		var NewLinkList = linkedList{
			Data:data,
			Next:point.Next,
		}
		point.Next = &NewLinkList
	}
	return true
}

// 返回值:成功/失败;被删除的元素
// 删除1:指定位置的节点
func (l *linkedList) Remove(index int) (bool,Elem) {
	if index < one || index > l.Len() {
		return false,-1
	} else {
		point := l
		for i := one; i < index; i++ {
			point = point.Next
		}
		data := point.Next.Data
		point.Next = point.Next.Next
		return true,data
	}
}
// 删除2:首个元素
func (l *linkedList) RemoveFirst() (bool,Elem) {
	if l.IsEmpty(){
		return false,-1
	}
	data := l.Next.Data
	l.Next = l.Next.Next
	return true,data
}
// 删除3:末尾元素
func (l *linkedList) RemoveLast() (bool,Elem) {
	if l.IsEmpty(){
		return false,-1
	}
	temp := l
	for temp.Next.Next != nil{
		temp = temp.Next
	}
	data := temp.Next.Data
	temp.Next = nil
	return true,data
}

// 修改1:首个元素(空表默认添加)
func (l *linkedList) SetFirst(elem Elem) {
	if l.IsEmpty(){
		n := linkedList{Data:elem}
		l.Next = &n
	}
	l.Next.Data = elem
}
// 修改2:末尾元素(空表添加)
func (l *linkedList) SetLast(elem Elem) {
	if l.IsEmpty(){
		n := linkedList{Data:elem}
		l.Next = &n
	}
	temp := l.Next
	for temp.Next != nil{
		temp = temp.Next
	}
	temp.Data = elem
}
// 修改3:任意位置添加
func (l *linkedList) Set(index int,elem Elem) bool {
	if index < one || index > l.Len(){
		return false
	}
	temp := l.Next
	for i := one;i < index;i ++ {
		temp = temp.Next
	}
	temp.Data = elem
	return true
}

// 取值1:取首
func (l *linkedList) GetFirst() (bool,Elem) {
	if l.IsEmpty(){
		return false,-1
	}
	return true,l.Next.Data
}
// 取值2:取尾
func (l *linkedList) GetLast() (bool,Elem) {
	if l.IsEmpty(){
		return false, -1
	}
	temp := l.Next
	for temp.Next != nil{
		temp = temp.Next
	}
	return true, temp.Data
}
// 取值3:取任意位置
func (l *linkedList) Get(index int) (bool,Elem) {
	if index < one || index > l.Len(){
		return false, -1
	}
	temp := l.Next
	for i := one;i < index;i++{
		temp = temp.Next
	}
	return true, temp.Data
}
