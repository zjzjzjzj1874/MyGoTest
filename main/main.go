package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	for i := 0;i < 10000;i ++ {
		go appendValueNoMutex(i)	// 不加锁会丢失
		//go appendValueWithMutex(i)	// 加锁不会丢失
	}
	//time.Sleep(time.Second)
	sort.Ints(s)
	for i,v := range s{
		fmt.Println(i,":",v)
	}
}

var s []int
var lock sync.Mutex

// 不加锁
func appendValueNoMutex(i int)  {
	s = append(s, i)
}

func appendValueWithMutex(i int)  {
	lock.Lock()
	s = append(s, i)
	lock.Unlock()
}



