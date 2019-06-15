package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {

	//多线程处理
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	//创建了100个生产的协程,但是消费的协程只有一个
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done) //如何将这个也改造成多协程的啊
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done //channel管道中取值
	if d == false {
		fmt.Println("\r\n writing failed.")
	} else {
		fmt.Println("\r\n writing successfully.")
	}
}

// select和case的用法
// 可使用一个废弃 channel done 来告诉剩余的 goroutine 无需再向 ch 发送数据。此时 <- done 的结果是 {}：
func selectCase() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	fmt.Println("Result total: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}

//从0-999中生成随机数字
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n //往data的channel中添加值
	wg.Done()
}

//消费随机数的函数
func consume(data chan int, done chan bool) {
	file, err := os.Create("concurrent")
	if err != nil {
		fmt.Println("\r\n create file err == ", err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(file, d)
		if err != nil {
			fmt.Println("\r\n write err == ", err)
			file.Close()
			done <- false
			return
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println("\r\n close file err == ", err)
		done <- false
		return
	}
	done <- true
}

//最简单的go协程 -- 通过sleep来完成
func ThreadGoroutine(a time.Duration) {
	fmt.Println("In ThreadGoroutine()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(a * 1e9)
	/*我们让 main() 函数暂停 10 秒从而确定它会在另外两个协程之后结束。
	如果不这样（如果我们让 main() 函数停止 4 秒），main() 会提前结束，
	longWait() 则无法完成。如果我们不在 main() 中等待，协程会随着程序的结束而消亡。*/
	fmt.Println("At the end of main()")
}
func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

//go 关键词的测试
func goRuntineTest() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i1: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
