package dataStructure

import (
	"sort"
	"sync"
)

// set集合:确定性;互异性;无序性
// 给定元素,要么属于集合,要么不属于 -- 确定性;不能存储重复的元素 -- 互异性;Set集合没有索引 -- 无序性;

// go实现的int类型集合:线程安全的
type Set struct {
	data map[int]bool
	sync.RWMutex
}

// 初始化Set
func NewSet() *Set {
	return &Set{data: map[int]bool{}}
}

// 判断set是否为空
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// 向Set中添加元素:是否需要=>返回值=-1?失败:set长度
func (s *Set) Add(elem int) {
	s.Lock()
	defer s.Unlock()
	s.data[elem] = true
}

// 移除集合中元素
func (s *Set) Remove(elem int) {
	s.Lock()
	delete(s.data, elem)
	s.Unlock()
}

// 判断元素是否在set中
func (s *Set) Contains(elem int) bool {
	return s.data[elem]
}

// 转化为list
func (s *Set) List() []int {
	s.RLock()
	var arr []int
	for i := range s.data {
		arr = append(arr, i)
	}
	s.RUnlock()
	return arr
}

// 转化为有序集合 -- TODO 是否是安全的?
func (s *Set) SortList() (arr []int) {
	arr = s.List()
	sort.Ints(arr)
	return
}

// 集合的长度
func (s *Set) Len() int {
	return len(s.List())
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	s.data = nil
	s.Unlock()
}

// 集合的相关运算
// 两个集合的交集
func (s *Set) InterSection(o *Set) *Set {
	s.Lock()
	defer s.Unlock()
	if s.Len() == zero || o.Len() == zero{
		return NewSet()
	}
	nS := NewSet()
	for s1 := range s.data {
		for s2 := range o.data {
			if s1 == s2 {
				nS.Add(s1)
				break
			}
		}
	}
	return nS
}

// 两个集合的并集
func (s *Set) Union(o *Set) *Set {
	s.RLock()
	defer s.RUnlock()
	if s.Len() == zero && o.Len() == zero{
		return NewSet()
	}
	nS := NewSet()
	for s1 := range s.data {
		nS.data[s1] = true
	}
	for s2 := range o.data {
		nS.data[s2] = true
	}
	return nS
}

// 两个集合是否相等
func (s *Set) Equal(o *Set) bool {
	s.RLock()
	defer s.RUnlock()
	if s.Len() != o.Len() {
		return false
	}
	arr1 := s.SortList()
	arr2 := o.SortList()
	for i := zero; i < s.Len(); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// 判断s是否是motherSet的子集 -- 不判断真子集
// 空集是任何集合的子集
func (s *Set) IsSubSet(motherSet *Set) bool {
	s.RLock()
	defer s.RUnlock()
	if motherSet.Len() < s.Len(){
		return false
	}
	for v := range s.data{
		if !motherSet.Contains(v){
			return false
		}
	}
	return true
}

// 计算差集 difference set :就是把set1中属于set2的元素去掉
func (s *Set) DiffSet(s2 *Set) *Set {
	s.RLock()
	s.RUnlock()
	if s.Len() == zero || s2.Len() == zero{
		return s
	}
	ns := NewSet()
	for v1 := range s.data{
		if !s2.Contains(v1){
			ns.Add(v1)
		}
	}
	return ns
}

// 对称差集 symmetric difference set :空集和其他集合的对称差集
func (s *Set) SymDiffSet(s2 *Set) *Set {
	return s.DiffSet(s2).Union(s2.DiffSet(s))
}