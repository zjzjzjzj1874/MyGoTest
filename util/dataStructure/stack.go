// The stack of array implementation has been completed;
// todo :the stack of linked list implementation has not been completed yet.
package dataStructure

type (
	// golang 使用切片实现栈
	SliceStack struct {
		Elem []Element
	}
	// golang使用双向链表实现栈
	LinkedStack struct {
		Prev *LinkedStack
		Next *LinkedStack
		Elem Element
	}
)

// 这里使用接口定义stack类型
type Stack interface {
	Push(interface{}) // 入栈
	Size() int        // 大小
	Clear()           // 清空
	IsEmpty() bool    // 是否为空
	Top() Element     // 栈顶元素 不删除
	Pop()             // 出栈
}

// region todo 链表实现栈

func linkedStackConstructor(item interface{}) *LinkedStack {
	return &LinkedStack{Prev: nil, Next: nil, Elem: item}
}

// 初始化链表的头节点
func NewLinkedStack() *LinkedStack {
	header := linkedStackConstructor(nil)
	return header
}

func (s *LinkedStack) Push(item interface{}) {
	ls := linkedStackConstructor(item)
	ls.Prev = s
	s.Next = ls
}

func (s *LinkedStack) Size() int {
	return 0
}

func (s *LinkedStack) Clear() {

}

func (s *LinkedStack) IsEmpty() bool {
	return false
}

func (s *LinkedStack) Top() Element {
	return nil
}

func (s *LinkedStack) Pop() {

}

// endregion 链表实现栈

// region 切片实现栈
// 新建切片Stack
func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Size() int {
	return len(s.Elem)
}

func (s *SliceStack) Clear() {
	s.Elem = nil
}

func (s *SliceStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceStack) Top() Element {
	if s.IsEmpty() {
		return nil
	}
	return s.Elem[s.Size()-1]
}

func (s *SliceStack) Push(item interface{}) {
	s.Elem = append(s.Elem, item)
}

func (s *SliceStack) Pop() {
	if !s.IsEmpty() {
		s.Elem = append(s.Elem[:s.Size()-1])
	}
}

// endregion 切片实现栈
