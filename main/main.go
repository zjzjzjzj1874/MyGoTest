package main

import (
	"encoding/json"
	"fileOperation"
	"fmt"
	"lib"
	"runtime"
	"sync"
	"time"
)

type Person struct {
	Name        string `json:"name"`          // json:"name" 表示在json中的key == name;如果没有的话,json中的key = Name
	Age         int64  `json:"age,omitempty"` // omitempty 忽略空值:如果没有该关键字,key为空的时候也会显示出来;有的话,key为空不会出现该字段的相关信息
	Tel         string `json:"tel,omitempty"`
	Pwd1        string `json:"-"`       // json:"-" 表示忽略这个字段;(对于密码这种字段起到保护的作用,不输出)
	Pwd2        string `json:"-,"`      // json:"-," 表示不会忽略这个字段;并且key显示为 "-"
	Int64String int64  `json:",string"` // json:",string" 将该类型转化为string;
	// Field is ignored by this package.
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

type Work struct {
}

func init() {
	fmt.Println("\r\n do in first init111")
}
func init() {
	fmt.Println("\r\n do in first init222")
}

func main() {

	//testArrayPoint()
	//one := 1
	//one,two := 1,2

	//ThreadGoroutine(10)

	//all := []map[string]interface{}{
	//	{
	//		"category_id": 41,
	//		"category_name": "category111-en",
	//		"category_name_alias": "{\"40\":\"category111-en\"}",
	//		"category_parent_id": 0,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 42,
	//		"category_name": "category222-food",
	//		"category_name_alias": "{\"40\":\"category222-food\"}",
	//		"category_parent_id": 0,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 46,
	//		"category_name": "category333-en",
	//		"category_name_alias": "{\"40\":\"category333-en\"}",
	//		"category_parent_id": 0,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 47,
	//		"category_name": "333--数码",
	//		"category_name_alias": "{\"40\":\"333--数码\"}",
	//		"category_parent_id": 46,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 48,
	//		"category_name": "Digital",
	//		"category_name_alias": "{\"55\":\"Digital\",\"56\":\"数码\",\"57\":\"デジタル\",\"58\":\"цифровой\",\"59\":\"Digital\"}",
	//		"category_parent_id": 47,
	//		"is_shown_index": "off",
	//	},
	//	{
	//		"category_id": 49,
	//		"category_name": "11-en墨镜",
	//		"category_name_alias": "{\"40\":\"11-en墨镜\"}",
	//		"category_parent_id": 41,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 50,
	//		"category_name": "22-化妆刷",
	//		"category_name_alias": "{\"40\":\"22-化妆刷\"}",
	//		"category_parent_id": 42,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 52,
	//		"category_name": "44-碧螺春",
	//		"category_name_alias": "{\"40\":\"44-碧螺春\"}",
	//		"category_parent_id": 49,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 59,
	//		"category_name": "11-墨--test",
	//		"category_name_alias": "{\"40\":\"11-墨--test\"}",
	//		"category_parent_id": 49,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 69,
	//		"category_name": "22-面粉",
	//		"category_name_alias": "{\"40\":\"22-面粉\"}",
	//		"category_parent_id": 41,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 71,
	//		"category_name": "墨镜测试分类 333",
	//		"category_name_alias": "{\"40\":\"墨镜测试分类 333\"}",
	//		"category_parent_id": 49,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 72,
	//		"category_name": "墨镜444",
	//		"category_name_alias": "{\"40\":\"墨镜444\"}",
	//		"category_parent_id": 49,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 73,
	//		"category_name": "4444",
	//		"category_name_alias": "{\"40\":\"4444\",\"41\":\"\"}",
	//		"category_parent_id": 0,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 74,
	//		"category_name": "墨镜444儿子",
	//		"category_name_alias": "{\"40\":\"墨镜444儿子\",\"41\":\"墨镜444儿子\"}",
	//		"category_parent_id": 72,
	//		"is_shown_index": "on",
	//	},
	//	{
	//		"category_id": 75,
	//		"category_name": "444孙",
	//		"category_name_alias": "{\"40\":\"444孙\",\"41\":\"孙子\"}",
	//		"category_parent_id": 74,
	//		"is_shown_index": "on",
	//	},
	//}
	//
	//topId := lib.NewSet()
	//for _,v := range all{
	//	if lib.InterfaceToString(v["category_parent_id"]) == "0"{
	//		id := lib.StringToInt(lib.InterfaceToString(v["category_id"]))
	//		topId.Add(id)
	//	}
	//}
	//fmt.Println("\r\n top id === ",topId)

	//pwd1 := lib.Md5Encode("zhoujian")
	//pwd2 := util.Md5Encode("zhoujian")
	//
	//fmt.Println("\r\n pwd1 ====== ",pwd1)
	//fmt.Println("\r\n pwd2 ====== ",pwd2)
	//
	//var house = "A B C"
	//ptr := &house
	//fmt.Println("\r\n ptr ==== ",ptr)
	//
	//str := "A B C"
	//var ptr1 *string
	//ptr1 = &str
	//fmt.Println("\r\n ptr1 ==== ",ptr1)
	//
	//m2 := map[string]interface{}{
	//	"name":"jack",
	//	"age":18,
	//}
	//ptr2 := &m2
	//m3 := map[string]interface{}{
	//	"name":"jack",
	//	"age":18,
	//}
	//ptr3 := &m3
	//fmt.Println("\r\n ptr2 ==== ",ptr2)
	//fmt.Println("\r\n ptr3 ==== ",ptr3)
	//fmt.Println("\r\n m2 ==== ",m2)
	//fmt.Println("\r\n m3 ==== ",m3)

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

//test fibonaci
func testFibonaci() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonaci(i))
	}
}

//go实现斐波那契数列
func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}

//递归调用函数求阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

// 数组指针和指针数组;
// 取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
// 一句话总结：注意*与谁结合，如p *[5]int，*与数组结合说明是数组的指针；
// 如p [5]*int，*与int结合，说明这个数组都是int类型的指针，是指针数组(由指针构成的数组) ---  p2 ====  [0xc000058400 0xc000058408 <nil> <nil> <nil>] -- 指针数组。
// PS:数组的指针 -- 只有一个地址,且改地址存放该数组;指针数组 -- 对应的非空元素都有一个地址
func testArrayPoint() {

	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)
	fmt.Printf("ip 变量的地址: %x\n", &ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	//数组指针,
	aa := [...]int{1, 2, 3, 4, 5} // -- 定义这个a,只是为了给数组指针赋值
	var p *[5]int = &aa
	fmt.Println(p, &p, *p)
	for i, v := range p {
		fmt.Println("\r\n p === ", i, v)
	}
	for i, v := range *p {
		fmt.Println("\r\n *p ==== ", i, v)
	}
	//指针数组
	var p2 [5]*interface{}
	p2[0] = setPointValue(0)
	p2[0] = setPointValue("false")
	p2[1] = setPointValue(2)
	fmt.Println("\r\n p2 ==== ", p2)
	fmt.Println("\r\n &p2 ==== ", &p2)
	for i, v := range p2 {
		if v != nil { // --添加这个判断:因为空指针是不能访问的
			fmt.Println("\r\n i,*v ====== ", i, *v)
			fmt.Println("\r\n i,&v ====== ", i, &v)
		} else {
			fmt.Println("\r\n i,v ====== ", i, v)
		}
	}

}

//给指针赋值
func setPointValue(i interface{}) *interface{} {
	return &i
}

//测试interface去重;
func testRemoveRepeatedMap() {
	var map1 []map[string]interface{}
	m1 := map[string]interface{}{
		"goods_id":     1,
		"goods_sku_id": 1,
	}
	m2 := map[string]interface{}{
		"goods_id":     2,
		"goods_sku_id": 2,
	}
	m3 := map[string]interface{}{
		"goods_id":     3,
		"goods_sku_id": 2,
	}
	map1 = append(map1, m1)
	map1 = append(map1, m1)
	map1 = append(map1, m1)
	map1 = append(map1, m2)
	map1 = append(map1, m3)
	map1 = append(map1, m3)
	fmt.Println("\r\n map111 ==== ", map1)
	map2 := RemoveRepeatedMap(map1)
	//RemoveRepeatMap(map1)
	fmt.Println("\r\n map222 ==== ", map2)
}

// 作用:map[string]interface{}数组去重
// 原理:定义一个新切片（数组），存放原数组的第一个元素，然后将新切片（数组）与原切片（数组）的元素一一对比，如果不同则存放在新切片（数组）中
// 参数:goods_id和goods_sku_id是需要比较的两个参数,可以减少也可以扩展
func RemoveRepeatedMap(map1 []map[string]interface{}) (newMap []map[string]interface{}) {
	newMap = make([]map[string]interface{}, 0)
	for i := 0; i < len(map1); i++ {
		repeat := false
		for j := i + 1; j < len(map1); j++ {
			m1 := map1[i]
			m2 := map1[j]
			if lib.InterfaceToString(m1["goods_id"]) == lib.InterfaceToString(m2["goods_id"]) && lib.InterfaceToString(m1["goods_sku_id"]) == lib.InterfaceToString(m2["goods_sku_id"]) {
				repeat = true
				break
			}
		}
		if !repeat {
			newMap = append(newMap, map1[i])
		}
	}
	return
}

//string数组去重,int数组也同理
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//数组合并,并简单处理
func returnHistoryListString(list, ids []string) (newList string) {
	newList = ""
	list = append(list, ids...)
	list = RemoveRepeatedElement(list)
	fmt.Println("\r\n 新的list == ", list)
	if len(list) <= 6 {
		for index, val := range list {
			if index == 0 {
				newList = val
			} else {
				newList += "-" + val
			}
		}
		fmt.Println("\r\n length <= 6,NewList === ", newList)
	} else {
		for index, val := range list {
			if index == len(list)-6 {
				newList = val
			} else if index > len(list)-6 {
				newList += "-" + val
			}
		}
		fmt.Println("\r\n length > 6,NewList === ", newList)
	}
	return
}

//读取控制台输入
func GetConsoleInput() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

//熟悉结构体的定义与修改
//结构体的三种定义方式
func StructTest() {
	//定义1
	p1 := new(Person)
	p1.Name = "张三"
	p1.Age = 18
	p1.Tel = "13344556677"
	//定义2
	p2 := Person{"李四", 17, "li111111", "admin", "123456", 64}
	//定义3
	p3 := Person{Name: "王二", Age: 20}
	fmt.Println("\r\n p1 =========== ", p1)
	fmt.Println("\r\n p2 =========== ", p2)
	fmt.Println("\r\n p3 =========== ", p3)
}

//JS中的json.parse方法:将json的字符串变为json类型;go中对应json.Unmarshal()方法
func JsonParse(str string) map[string]interface{} {
	person := map[string]interface{}{}
	error := json.Unmarshal([]byte(str), &person)
	if error != nil {
		fmt.Println("\r\n error === ", error)
		return person
	}
	return person
}

//JS中的json.stringy方法:将json类型变为json字符串;go中对应json.Marshal()方法
func JsonStringy(TestMap interface{}) (string, bool) {
	str, err := json.Marshal(TestMap)
	if err != nil {
		fmt.Println("\r\n err ==== ", err)
		return err.Error(), false
	}
	return string(str), true
}
