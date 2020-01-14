package dataStructure

import "sync"

// golang实现队列

type (
	Queue struct {
		Item []Element
		lock sync.RWMutex // 读写锁
	}
)

// 创建队列
func NewQueue() *Queue {
	return &Queue{}
}
