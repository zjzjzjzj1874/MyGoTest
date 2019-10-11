package main

import (
	"fmt"
	"math/big"
	"os/user"
	"reflect"
	"strconv"
	"sync"
	"time"
)

// 下面两个函数是有区别的
func errTest(arr [3]int) error {
	return nil
}
func errTest1(arr []int) error {
	return nil
}
func add(a []int, b int, c int) (d []int) {
	arr1 := new([]int)

	*arr1 = append(*arr1, a[:b]...)

	*arr1 = append(*arr1, c)

	*arr1 = append(*arr1, a[b:]...)

	fmt.Println(a)
	fmt.Println(*arr1)
	return a
}

func add1(a []int, b int, c int) (d []int) {
	var abcd []int

	abcd = append(abcd, a[:b]...)

	abcd = append(abcd, c)

	abcd = append(abcd, a[b:]...)

	fmt.Println(abcd)
	fmt.Println(a)
	return a
}

func add2(a []int, b int, c int) (d []int) {

	fmt.Println("before append:a == >", a)

	abcd := append(a[:b], c)

	fmt.Println("after append:a == >", a)

	fmt.Println("after append:abcd == >", abcd)
	return a
}

// 指针类型
func pti() {
	var j int = 5
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(&j))
	fmt.Println(j)
	fmt.Println(&j)
	i := new(int)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(*i))
	fmt.Println(reflect.TypeOf(&i))
	fmt.Println(i)
	fmt.Println(*i)
	fmt.Println(&i)
}

//var ARGS string
//
//func main() {
//
//	var uptime *bool = new(bool)
//	flag.BoolVar(uptime,"u", false, "print system uptime")
//	flag.Parse()
//
//	ARGS = strings.Join(flag.Args(), " ")
//	if len(os.Args) < 2 {
//		flag.Usage()
//		os.Exit(1)
//	}
//
//	if *uptime {
//		fmt.Println("12 days")
//	}
//}

func main() {

	//l := list.New()
	//l.PushFront(1)
	//l.PushFront(2)
	//l.PushFront(3)
	//l.PushFront(4)
	//l.PushFront(4)
	////l.PushFront(5)
	//
	//fmt.Println(l.Len())
	//for e := l.Front(); e != nil; e = e.Next() {
	//	if i, ok := e.Value.(int); ok {
	//		fmt.Println(e.Value,";i = ",i)
	//	}
	//}
	//fmt.Println()

	u, err := user.Current()
	if err != nil {
		fmt.Printf("Current: %v (got %#v) \n", err, u)
	}
	if u.HomeDir == "" {
		fmt.Printf("didn't get a HomeDir")
	}
	fmt.Println(u.HomeDir)
	if u.Username == "" {
		fmt.Printf("didn't get a username")
	}
	fmt.Println(u.Username)
	//pti()

	//abc := []int{1, 2, 3, 4, 5, 6}
	//var x int = 3
	//var i int = 100
	//add2(abc, x, i)

	//s := ""
	//fmt.Println(s)
	//fmt.Println(&s)
	//
	//err := errTest()
	//fmt.Println(err)
	//fmt.Println(&err)

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
