package _chan

import "fmt"

type (
	Person struct {
		Name string
		Age  string
		Addr Addr
	}
	Addr struct {
		City string
	}
)

func TranChan() {
	pc := make(chan Person, 2)
	p1 := Person{Name: "XiaoMing", Age: "18"}
	pc <- p1
	p1.Addr.City = "成都"
	fmt.Println("person 1 :", p1)
	fmt.Println("person chan:", pc)

	p2 := Person{Name: "张三", Age: "17"}
	pc <- p2 // 将p2写入person channel中,即生产channel
	fmt.Println("person chan:", pc)

	go func() {
		for {
			select {
			case v := <-pc: // 读取channel=>即消费channel
				fmt.Println("v in person chan = ", v)
			}
		}
	}()
	//np := <- pc	// 读取channel中数据,每读一次,少一条
	//fmt.Println("new person = ",np)
}
