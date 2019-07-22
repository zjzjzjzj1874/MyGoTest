package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

func main() {


	fmt.Println(ToBigFloat("12.df2354541bhj"))
	fmt.Println(BigFloatToString(big.NewFloat(10.241)))
	fmt.Println(BigFloatToString(big.NewFloat(10.241254)))
	fmt.Println(BigFloatToString(big.NewFloat(10.2419254)))
	fmt.Println(BigFloatToString(big.NewFloat(10.24)))
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
	f, _, _ := big.ParseFloat(str, 10, 12, big.ToNearestEven)
	return f
}

// big.float转string(保留三位小数)
func BigFloatToString(b *big.Float) string {
	return b.Text('f',3)
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f,'f',-1,64)
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



