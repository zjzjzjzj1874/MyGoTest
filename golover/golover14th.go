package golover

import (
	"context"
	"fmt"
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

func closure() {
	arr := []int{1, 2, 3, 4, 5, 6}
	for i := range arr {
		go func() {
			fmt.Println("闭包无参数==>", arr[i])
		}()
	}
	time.Sleep(300 * time.Millisecond) // 如果不等待,用go开启的线程可能会因为主程序退出,而不会全部执行

	for i := range arr {
		go func(j int) {
			fmt.Println("闭包传参==>", arr[j])
		}(i)
	}
	time.Sleep(300 * time.Millisecond)

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
// endregion
