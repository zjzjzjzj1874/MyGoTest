package dataStructure

// TODO 待完成:
//  1.线性表中存储数据不仅限于int

const(
	doubleTippingPoint 	= 100
	double				= 2
)

const(
	zero = iota
	one
	two
)

type (
	// 线性表中的元素
	LinearElem int

	// 定义一个线性表结构体
	LinearList struct {
		Maxsize int          // 最大长度 TODO 是否有必要,可以让其 = len(LinearList.Data) =>更安全;别人可能把Maxsize修改了,但是Data却还是原来的容量
		CurLen  int          // 当前长度
		Data    []LinearElem // 保存数据(也可以把linearElem换成interface{},更加通用)
	}
)

// 初始化一个线性表
func InitLinearList(maxsize int) *LinearList {
	if maxsize < one{
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
	//return l.Maxsize == l.CurLen
	return len(l.Data) == l.CurLen
}

// 判断索引是否越界
func (l *LinearList) IndexOutOfRange(index int) bool {
	//return index > l.Maxsize || index < 1
	return index > len(l.Data) || index < one
}

// 根据index获取对应元素
func (l *LinearList) GetElem(index int) (LinearElem, bool) {
	if l.IndexOutOfRange(index) {
		return -one, false
	}
	return l.Data[index - one], true
}

// 插入元素
// 1.末尾append一个数据 当前长度++
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) Append(ele LinearElem) int {
	if l.IsFull() {
		return  -one
	}
	l.Data[l.CurLen] = ele
	l.CurLen++
	return l.CurLen
}
// 2.在index位置前插入元素e;
// 返回值 -1失败;成功后返回:插入后的线性表长度
func (l *LinearList) Insert(index int, e LinearElem) int {
	if l.IsFull() || index < one || index > l.CurLen  {
		return -one
	}
	for i := l.CurLen;i > index - one;i--{
		l.Data[i] = l.Data[i - one]
	}
	l.Data[index - one] = e
	l.CurLen ++
	return l.CurLen
}

// 删除元素
// 1.删除最后一个元素
// 返回值 -1失败;成功后返回:删除后的线性表长度
func (l *LinearList) DelLast() int {
	if l.IsEmpty(){
		return -one
	}
	l.Data[l.CurLen - one] = 0
	l.CurLen --
	return l.CurLen
}
// 2.删除指定位置元素
// 返回值:-1失败  成功后返回:删除后的线性表长度
func (l *LinearList) Delete(index int) int {
	if l.IsEmpty() || index > l.CurLen || index < one{
		return -one
	}
	for i := index;i < l.CurLen; i++{
		l.Data[i-one] = l.Data[i]
	}
	l.Data[l.CurLen-one] = 0
	l.CurLen --
	return l.CurLen
}

// 修改替换
func (l *LinearList) Update(index int,ele LinearElem ) int {
	if  index > l.CurLen || index < one{
		return -one
	}
	l.Data[index - one] = ele
	return index
}

// 扩容:i<1的时候按规则扩容,指定长度的话,按照指定长度扩容
// 扩容规则:总长度小于100的时候,扩容两倍;其他扩容1.5倍
func (l *LinearList) Extend(i int) *LinearList {
	maxsize := len(l.Data)
	if i < one && maxsize < doubleTippingPoint{
		maxsize *= double
	}else if i < one && maxsize >= doubleTippingPoint {
		maxsize += (maxsize%two + maxsize)/two
	}else{
		maxsize += i
	}
	extend := make([]LinearElem,maxsize - len(l.Data))
	l.Data = append(l.Data, extend...)
	l.Maxsize = len(l.Data)
	return l
}

// 两个线性表的合并
func (l *LinearList) Combine(l2 *LinearList) *LinearList {
	l2Len := make([]LinearElem,len(l2.Data))
	l.Data = append(l.Data, l2Len...)
	l.Maxsize = len(l.Data) + len(l2.Data)
	for i := zero;i < l2.CurLen;i ++ {
		l.Append(l2.Data[i])
	}
	return l
}