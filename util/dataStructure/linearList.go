package dataStructure

// TODO 待完成:
//  1.两个线性表的合并
//  2.线性表中存储数据不仅限于int
type (
	// 线性表中的元素
	LinearElem int

	// 定义一个线性表结构体
	LinearList struct {
		Maxsize int          // 最大长度
		CurLen  int          // 当前长度
		Data    []LinearElem // 保存数据(也可以把linearElem换成interface{},更加通用)
	}
)

// 初始化一个线性表
func InitLinearList(maxsize int) *LinearList {
	return &LinearList{Maxsize: maxsize, Data: make([]LinearElem, maxsize)}
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
func (l *LinearList) GetElem(index int) (LinearElem, bool) {
	if l.IndexOutOfRange(index) {
		return -1, false
	}
	return l.Data[index-1], true
}

// 插入元素
// 1.末尾append一个数据 当前长度++
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) Append(ele LinearElem) int {
	if l.IsFull() {
		return  -1
	}
	l.Data[l.CurLen] = ele
	l.CurLen++
	return l.CurLen
}
// 2.在index位置前插入元素e;
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) Insert(index int, e LinearElem) int {
	if l.IsFull() || index < 1 || index > l.CurLen  {
		return -1
	}
	for i := l.CurLen;i > index - 1;i--{
		l.Data[i] = l.Data[i-1]
	}
	l.Data[index-1] = e
	l.CurLen ++
	return l.CurLen
}

// 删除元素
// 1.删除最后一个元素
// 返回值 -1失败;成功后返回:删除后的线性表长度
func (l *LinearList) DelLast() int {
	if l.IsEmpty(){
		return -1
	}
	l.Data[l.CurLen - 1] = 0
	l.CurLen --
	return l.CurLen
}
// 2.删除指定位置元素
// 返回值:-1失败  成功后返回:删除后的线性表长度
func (l *LinearList) Delete(index int) int {
	if l.IsEmpty() || index > l.CurLen || index < 1{
		return -1
	}
	for i := index;i < l.CurLen; i++{
		l.Data[i-1] = l.Data[i]
	}
	l.Data[l.CurLen-1] = 0
	l.CurLen --
	return l.CurLen
}

// 修改替换
func (l *LinearList) Update(index int,ele LinearElem ) int {
	if  index > l.CurLen || index < 1{
		return -1
	}
	l.Data[index - 1] = ele
	return index
}