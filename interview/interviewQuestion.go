package interview

import (
	"fmt"
	"runtime"
	"sync"
)

//dafer 1
func DeferCall() {
	defer func() { fmt.Println(" 打印前") }()
	fmt.Println("前")
	defer func() { fmt.Println(" 打印中") }()
	fmt.Println("中")
	defer func() { fmt.Println(" 打印后") }()
	fmt.Println("后")

	//输出结果:(没有defer的话正常走,有defer的都是先进后出)
	//前 中 后  打印后 打印中 打印前
}

//defer 2
func Defer222() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	//解析:这道题类似第1题 需要注意到defer执行顺序和  --值传递:虽然函数不执行,但是参数的值已经赋予了.
	// index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用calc("10",1,2)==>	10,1,2,3
	// 执行index:2时,与之前一样，需要先调用calc("20",0,2)==>									20,0,2,2
	// 执行到b=1时候开始调用，index:2==>calc("2",0,2)==>									2,0,2,2
	// 最后执行index:1==>calc("1",1,3)==>													1,1,3,4
}

// 计算
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type student struct {
	Name string
	Age  int
}

//指针问题:避免重复分配内存 -- 但是不是很友好
func Pase_student() {
	m := make(map[string]*student)
	stus1 := []*student{
		&student{Name: "zhou", Age: 24},
		&student{Name: "li", Age: 23},
		&student{Name: "wang", Age: 22},
		&student{Name: "wang", Age: 24},
	}
	for _, stu := range stus1 {
		m[stu.Name] = stu
	}
	fmt.Println("\r\n m ======== ", m)
	for i, v := range m {
		fmt.Println(i)
		fmt.Println(*v)
	}
	fmt.Println("\r\n 正确示范:", stus1) //正确示范: [0xc0000046e0 0xc000004700 0xc000004720 0xc000004740]

	//错误赋值示范 start
	stus2 := []student{
		{Name: "zhou", Age: 16},
		{Name: "li", Age: 17},
		{Name: "wang", Age: 18},
	}
	fmt.Println("\r\n 错误示范:", stus2) //错误示范: [{zhou 24} {li 23} {wang 22}]  即使内容不同,但三个student的指针是指向同意不hash地址;
	for _, stu1 := range stus2 {
		m[stu1.Name] = &stu1
	}
	fmt.Println("\r\n m ======== ", m)
	for i, v := range m {
		fmt.Println(i)
		fmt.Println(*v)
	}
	//错误示范 end

	//或者这样也可以
	stus3 := []student{
		{Name: "zhou", Age: 16},
		{Name: "li", Age: 17},
		{Name: "wang", Age: 18},
	}
	for i := 0; i < len(stus3); i++ {
		m[stus3[i].Name] = &stus3[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

//go的随机性和闭包
func BiBaoHanShu() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i111: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i222: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 测试make和append : make是动态的
// 案例一
func AppendMake1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //打印的值:0 0 0 0 0 1 2 3
}

// 案例二
func AppendMake2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //打印的值:1 2 3
}
