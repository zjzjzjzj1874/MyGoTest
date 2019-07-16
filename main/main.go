package main

import (
	"fmt"
	"math/big"
	"sync"
)

func main() {

	a1 := ToBigFloat("100.234aj123")
	fmt.Println(a1)
	a2 := big.NewFloat(-100)


	fmt.Printf("sum ===  %s\n", a1.Add(a1,a2).String())
	//fmt.Printf("a2 = %s\n", a2.String())

	//for i := 0;i < 10000;i ++ {
	//	go appendValueNoMutex(i)	// 不加锁会丢失
	//	//go appendValueWithMutex(i)	// 加锁不会丢失
	//}
	////time.Sleep(time.Second)
	//sort.Ints(s)
	//for i,v := range s{
	//	fmt.Println(i,":",v)
	//}
}
func ToBigFloat(str string) *big.Float {
	f, _, _ := big.ParseFloat(str, 10, 256, big.ToNearestEven)
	return f
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



