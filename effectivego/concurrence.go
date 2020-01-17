package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex

func main() {
	withoutPointer()
	//withPointer()
}

func withPointer() {
	a := 1
	p_a := &a
	var wg sync.WaitGroup
	wg.Add(2)
	go func(p_a *int) {
		for i := 1; i < 100000; i++ {
			l.Lock()
			*p_a += 1
			l.Unlock()
		}
		wg.Done()
	}(p_a)
	go func(p_a *int) {
		for i := 1; i < 100000; i++ {
			l.Lock()
			*p_a -= 1
			l.Unlock()
		}
		wg.Done()
	}(p_a)
	wg.Wait()
	fmt.Println(*p_a)
}

func withoutPointer() {
	a := 1
	var wg sync.WaitGroup
	wg.Add(2)
	go func(a int) {
		for i := 1; i < 100000; i++ {
			l.Lock()
			a += 1
			l.Unlock()
		}
		wg.Done()
	}(a)
	go func(p_a int) {
		for i := 1; i < 100000; i++ {
			l.Lock()
			p_a -= 1
			l.Unlock()
		}
		wg.Done()
	}(a)
	wg.Wait()
	fmt.Println(a)
}
