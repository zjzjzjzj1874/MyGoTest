package dataStructure

type (
	// 线性表中的元素
	LinearElem int

	// 定义一个线性表结构体
	LinearList struct{
		Maxsize int				// 最大长度
		CurLen	int				// 当前长度
		Data	[]LinearElem	// 保存数据(也可以把linearElem换成interface{},更加通用)
	}
)

// 初始化一个线性表
func InitLinearList(maxsize int) *LinearList {
	return &LinearList{Maxsize:maxsize,Data:make([]LinearElem,maxsize)}
}

// 判断线性表是否为空
func (l *LinearList) IsEmpty() bool {
	return 0 == l.CurLen
}

// 判断线性表是否已满
func (l *LinearList) IsFull() bool {
	return l.Maxsize == l.CurLen
}

// 判断索引是否越界
func (l *LinearList) IndexOutOfRange(index int) bool {
	return index > l.Maxsize || index < 1
}

// 根据index获取对应元素
func (l *LinearList) GetElem(index int) (LinearElem,bool) {
	if l.IndexOutOfRange(index){
		return -1, false
	}
	return l.Data[index - 1], true
}

// TODO 插入元素 删除元素
// https://www.jianshu.com/p/68411ca93b67    https://studygolang.com/articles/11689
// 末尾　ｐｏｐ　一个数据
// 末尾　Append 一个数据
// 任意位置　ｉｎｓｅｒｔ　一个数据