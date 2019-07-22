// 简单工具类
package util

import (
	"encoding/json"
	"fmt"
)

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

//读取控制台输入
func GetConsoleInput() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
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

// nil or empty
type I interface {
	m()
}
type T struct {
	num int
}

//T类型的一个方法,空方法
func (t T) m() {
	fmt.Println("I’m alive!")
}

type P *T

//测试interface的nil和空
func Combine() {
	var i I //1:申明 type为I的变量:i
	i = T{} //2:赋值	i = 空的T {0}

	//上面两步可以一步执行;
	i1 := I(T{})
	i1.m()

	i.m()
}

//测试nil和empty的方法
func TestNilOrEmpty() {
	p := T{}           //这样申明:表示 p = T{0}  是一个空的T类型
	fmt.Println(p)     //打印值:{0}
	fmt.Println(p.num) //打印值:0
	p.m()              //没有问题,不会报错

	var p1 *T //这样申明:是一个nil的T,因为没有初始化
	p1 = &p
	fmt.Println(p1, ";address p1 == ", &p1)
	fmt.Println(p, ";address   p == ", &p)
	fmt.Println(p, ";p == ", *p1)
	p.m() //没问题,不报错
	//fmt.Println(p1.num)	//报错:p1是nil的,不存在num的属性;   ---- panic: runtime error: invalid memory address or nil pointer dereference

}
