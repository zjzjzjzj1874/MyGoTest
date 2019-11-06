package main

import (
	"fmt"
	"sync"
)

// Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs

func main() {
	//nilMap()
	waitGroupUse()
}

// region Compare nil map and empty map
// go中的map类似Java中的HashMap,都是线程不安全的.
func nilMap() {
	// 1.var初始化的map不能赋值 ==> 生成了一个nil map
	var m map[string]int // 生成一个nil map
	//m["one"] = 1 // error ==>assignment to entry in nil map
	fmt.Println("m === ", m, m, "m是否为nil=>", m == nil)
	fmt.Println("m的长度==", len(m))
	delete(m, "one ♥")

	// 2.make的初始化可以赋值 ==> 生成了一个空map
	m2 := make(map[string]int) // 生成一个空map
	fmt.Println("m2 是否等于nil:", m2 == nil)
	fmt.Println("m2的长度", len(m2))
	m2["one"] = 1
	fmt.Println("m2 === ", m2, &m2)

	// 3.初始化赋值一起
	m1 := map[string]int{
		"one": 1,
	}
	m1["one"] = 2
	m1["two"] = 1
	fmt.Println(m1)
}

// endregion using nil maps

// region goroutine中waitGroup的使用
func waitGroupUse() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		//go doit(i,wq,done,wg)	// 注意:这里如果传wg的值,那将是值的副本,必须要传指针
		go doit(i, wq, done, &wg)
	}

	for i := 0; i < workerCount; i++ {
		wq <- i
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running \n", workerId)
	defer wg.Done()
	for {
		select {
		case m := <-wq:
			fmt.Printf("[%v] m => %v\n", workerId, m)
		case <-done:
			fmt.Printf("[%v] is done \n", workerId)
			return
		}
	}
}

// endregion goroutine中waitGroup的使用
