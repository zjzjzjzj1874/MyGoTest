package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"
)

func errTest() error {
	return nil
}

func main() {

	s := ""
	fmt.Println(s)
	fmt.Println(&s)

	err := errTest()
	fmt.Println(err)
	fmt.Println(&err)

	//for {
	//	fmt.Println("请输入两个参数,用空格号隔开: ")
	//	fmt.Scanln(&firstName, &lastName)
	//	fmt.Printf("input == %v,%v \n", firstName, lastName)
	//}
	//fmt.Sscanf(firstName, lastName,  test1)
	//fmt.Println("From the string we read: ", firstName, lastName, test1)
	//fmt.Println("From the string we read: ", f, i, s)

	//fmt.Println(ToBigFloat("12.df2354541bhj"))
	//fmt.Println(BigFloatToString(big.NewFloat(10.241)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.241254)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.2419254)))
	//fmt.Println(BigFloatToString(big.NewFloat(10.24)))
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
	return b.Text('f', 3)
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

var s []int
var lock sync.Mutex

// 不加锁
func appendValueNoMutex(i int) {
	s = append(s, i)
}

func appendValueWithMutex(i int) {
	lock.Lock()
	s = append(s, i)
	lock.Unlock()
}

// 打招呼
func greet() {
	now := time.Now()
	zeroOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(zeroOfToday.Add(15 * time.Hour)) {

	}
}
