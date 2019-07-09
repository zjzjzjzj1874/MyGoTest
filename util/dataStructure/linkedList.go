package dataStructure

// go的单链表实现
type LinkedList struct{
	Data 	int
	Next	*LinkedList
}

// 初始化链表
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// 判断空
func (l *LinkedList) IsEmpty() bool {
	return l.Next == nil
}
