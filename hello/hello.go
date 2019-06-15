package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//这里说明：因为一个包里可以有多个init函数，执行顺序好像是从上到下的
func init() {
	//fmt.Println("\r\n hello from init11111")
}
func init() {
	//fmt.Println("\r\n hello from init22222")
}
func main() {
	ok, isExist, oldContent, msg := WriteStringToFile("test", "hello,it's me.")
	fmt.Println("\r\n ok ==== ", ok)
	fmt.Println("\r\n isExist ==== ", isExist)
	fmt.Println("\r\n oldContent ==== ", oldContent)
	fmt.Println("\r\n msg ==== ", msg)
	//str := "abc"
	//s := make([]string,0)
	//s1 := make([]string,0)
	//s = append(s, str)
	//
	//
	//copy(s, s1)
	//sort.Strings(s)
	//
	//
	//fmt.Println("\r\n s1 ==== ",s1)
	//fmt.Println("\r\n s ==== ",s)

	//fmt.Println("\r\n hello,go")
	//fmt.Println("\r\n the result === ",string2.Reverse("hello"))
}

//检查文件是否存在
func fileExist(fileName string) bool {
	_, err := os.Lstat(fileName)
	return !os.IsNotExist(err)
}

//埃克森尔面试题目
func WriteStringToFile(fileName, content string) (ok, isExist bool, oldContent, msg string) {
	isExist = fileExist(fileName)
	if len(content) < 8 {
		return false, isExist, "", "The length of the content is ought to more than 8."
	}
	if isExist {
		//文件存在，读取后重新写入
		dataBefore, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("read file occurred an err: " + err.Error())
		}
		if string(dataBefore) == "" {
			msg = "The content of old file is null."
		} else {
			oldContent = string(dataBefore)
		}
	}
	data := []byte(content)
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return false, isExist, "", "write string error,and err == " + err.Error()
	}
	return true, isExist, oldContent, msg
}
