package thread

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//func TestAppend(t *testing.T) {
//	x := []string{"start"}
//
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		y := append(x, "hello", "world")
//		t.Log(cap(y), len(y))
//	}()
//	go func() {
//		defer wg.Done()
//		z := append(x, "goodbye", "bob")
//		t.Log(cap(z), len(z))
//	}()
//	wg.Wait()
//}

//func TestAppend2(t *testing.T) {
//	x := make([]string, 0, 6)
//
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		y := append(x, "hello", "world")
//		t.Log(len(y))
//	}()
//	go func() {
//		defer wg.Done()
//		z := append(x, "goodbye", "bob")
//		t.Log(len(z))
//	}()
//	wg.Wait()
//}

func TestAppend3(t *testing.T) {
	x := make([]string, 0, 6)
	x = append(x, "start")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := make([]string, 0, len(x)+2)
		y = append(y, x...)
		y = append(y, "hello", "world")
		t.Log(cap(y), len(y), y[0])
		fmt.Println(time.Now().UnixNano()," === ",y)
	}()
	go func() {
		defer wg.Done()
		z := make([]string, 0, len(x)+2)
		z = append(z, x...)
		z = append(z, "goodbye", "bob")
		t.Log(cap(z), len(z), z[0])
		fmt.Println(time.Now().UnixNano()," === ",z)
	}()
	wg.Wait()
}