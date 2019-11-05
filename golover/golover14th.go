package golover

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)

type DataService struct{}

func NewDataService() DataService {
	return DataService{}
}

// Status will tell us that our service is ok
func (DataService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return today's date
func (DataService) Get(ctx context.Context) (string, error) {
	return time.Now().Format("2006/01/02"), nil
}

// Validate will check if the date today's date
func (DataService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("2006/01/02", date)
	if err != nil {
		return false, err
	}
	return true, nil
}

// region For语句中的迭代变量与闭包函数

// 闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。
// 另一种表示方式为：一个闭包继承了函数所声明时的作用域。
// 这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁

func closure() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("追踪==>%s:%d", file, line)
	}

	where()
	start := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6}
	for i := range arr {
		go func() {
			fmt.Println("闭包无参数==>", arr[i])
		}()
	}
	where()
	time.Sleep(300 * time.Millisecond) // 如果不等待,用go开启的线程可能会因为主程序退出,而不会全部执行

	for i := range arr {
		go func(j int) {
			fmt.Println("闭包传参==>", arr[j])
		}(i)
	}
	time.Sleep(300 * time.Millisecond)
	where()
	// 返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	// 它不关心这些捕获了的变量和常量是否已经超出了作用域
	// 所以只有闭包还在使用它，这些变量就还会存在。
	f := f2()
	fmt.Println("f === ", f()) // 1 所以这个i是一直存在的,而且每次会+1
	fmt.Println("f === ", f()) // 4
	fmt.Println("f === ", f()) // 9
	fmt.Println("f === ", f()) // 16
	// 变量的生命周期不由它的作用域决定：f2 返回后，变量 i 仍然隐式的存在于f中。

	c := closureAdd(10)
	fmt.Println("计算值==>", c(1))
	fmt.Println("计算值==>", c(1))
	fmt.Println("计算值==>", c(1))
	fmt.Println("计算值==>", c(1))
	end := time.Now()
	fmt.Println("运行时间==>", end.Sub(start))
}

func f2() func() int {
	var i int // 申明之后,只要f2存在,i就会隐形存在
	return func() int {
		i++
		return i * i
	}
}

// 这里的i,无论是传入参数,还是申明参数,只要这个闭包函数还存在,那么这个闭包的生命周期就是变量的周期
func closureAdd(i int) func(int) int {
	//i := 10
	return func(num int) int {
		i++
		return i + num
	}
}

// endregion

// region Defer关键字的使用

func deferWithParam() {
	fmt.Println("返回前打印ret ==>", deferT1())

	i := 0
	defer fmt.Println(i) // 0==>其实在进入defer时,程序就已经预设好要执行什么了
	i++
	return
}

// 变量 `ret` 的值为 2，因为 `ret++` 是在执行 `return 1` 语句后发生的。即:defer在return后面执行
func deferT1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// endregion

// region 程序运行时间统计

func calRuntime() {
	start := time.Now()
	// do what you want to do
	time.Sleep(100 * time.Millisecond)
	fmt.Println("hello world")
	end := time.Now()
	fmt.Println("运行时间==>", end.Sub(start))
}

// endregion

// region

type PersonChannel struct {
	Name string
	Age  int
}

// 结构体的channel测试
func personChan() {
	personCh := make(chan PersonChannel, 3)
	person1 := PersonChannel{"大明", 18}
	personCh <- person1 // 向通道中添加数据

	person1.Name = "小明"
	fmt.Println(person1)

	// 说明:channel传输自定义的对象时,如果修改了一段的数据,另一端的数据没影响,即:channel传递后的数据是相互独立的
	newPerson := <-personCh // 读取通道中数据
	fmt.Println("new person =>", newPerson)
}

// 多线程的初始化和退出
// 1.关闭后的通道仍可以读取,但是不能写入了
// 2.不能读取空的通道
func learnChannel() {
	ch := make(chan int, 5) // 有数字是带缓冲的通道,超过容量后也会阻塞
	sign := make(chan int, 2)

	//v1,ok := <- ch	// 2.异常=> 不能读取空的channel

	go func() { // 这里为啥要开启一个go才不会出错啊
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		fmt.Println("the channel is closed")

		sign <- 0
	}()

	val := <-ch
	fmt.Println("读取一个值==>", val) // 1.通道即使关闭了,依旧可以读取

	go func() {
		for {
			i, ok := <-ch
			fmt.Printf("%d, %v \n", i, ok)

			if !ok {
				break
			}

			time.Sleep(time.Second * 2)
		}
		sign <- 1
	}()

	<-sign
	<-sign
}

// 退出多通道
func quitChannel() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println("message from g =>", v)
			case <-quit:
				fmt.Println("退出B")
				break
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("quit methods")
}

//endregion
