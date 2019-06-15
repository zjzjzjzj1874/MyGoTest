package fileOperation

import (
	f "fileManager"
	"fmt"
	"io/ioutil"
	"lib"
	"os"
)

func init() {
	fmt.Println("do something from init function.")
}

func main() {
	//writeStringIntoFile("we are writing string words into file.")
	//writeContentLineByLine("mjd/line",[]string{"Hello from go.","Nice to meet you.","See you again."})
	//appendContent("mjd/","my.txt","What is this in Chinese?","")

	readAll()

}

//使用ioutil包创建文件,并写入内容
func createFileByIOUtil() {
	err := ioutil.WriteFile("concurrent", []byte("content test"), 666)
	if err != nil {
		fmt.Println("\r\n write file by ioutil fail,err == ", err)
	}
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
// 读取
func readAll() {
	dir, err := ioutil.ReadDir("mjd")
	if err != nil {
		fmt.Println("\r\n read Dir err === ", err)
		return
	}
	for i, v := range dir {
		fmt.Printf("\r\n i == %d,v == %v", i, v.Name())
		fmt.Println("\r\n Sys = ", v.Sys())
	}
	fmt.Println("\r", dir)
}

//根据名字读取文件
func readFile(name string) {
	con, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("\r\n read file fail,err == ", err)
		return
	}
	fmt.Println(string(con))
}

//go写入文件内容 OS包中的方法
//步骤:	1.创建文件/文件夹
// 		2.将内容写入文件
// 		3.关闭文件
func writeStringIntoFile(content string) {
	file, err := os.Create("mjd/myFile.txt") //1.创建文件
	if err != nil {
		fmt.Printf("create file fail,err = %d", err)
		file.Close()
		return
	}
	//length,err := file.WriteString(content)				//2.写内容 -- 字符串写入
	//length,err := file.Write([]byte(content))				//2.写内容 -- 字节写入
	length, err := file.WriteAt([]byte(content), 4) //2.写内容 -- 字节写入 偏移量:4个空格
	if err != nil {
		fmt.Printf("write string into file fail,err = %d", err)
		file.Close()
		return
	}
	fmt.Printf("Written successfully,the content you written is %d byte.", length)
	err = file.Close() //3.关闭文件
	if err != nil {
		fmt.Printf("close the file fail,err = %d.", err)
	}
}

//字符串一行一行写入文件
func writeContentLineByLine(fileName string, fileContent []string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("create", fileName, " fail,err = ", err)
		return
	}
	for _, v := range fileContent {
		n, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Printf("write line by line fail,err = %d", err)
			return
		}
		fmt.Printf("write line by line success,length = %d", n)
	}
	fmt.Printf("success.")
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹 TODO 可以优化 -- 简化代码
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

//追加内容到文件里面,或新建
// mode == "append"是追加,没有这个条件就是新建文件(如果已经存在,就截断文件并覆盖)
// 1.判断该文件是否存在
// 2.不存在就新建 存在就写入
// 3.写入
func appendContent(path, name, content, mode string) {
	ok := isExist(path + name)
	if !ok {
		fmt.Println("The file is not exist.")
		os.Create(path + name)
	}
	fileMode := os.O_CREATE
	fmt.Println("\r\n mode === ", fileMode)
	if mode == "append" {
		fileMode = os.O_APPEND | os.O_WRONLY
	} else {
		fileMode = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
	}
	f, err := os.OpenFile(path+name, fileMode, 666)
	if err != nil {
		fmt.Println("\r\n open file err == ", err)
	}
	defer f.Close()
	if err != nil {
		fmt.Println("\r\n write string err == ", err)
	}
}
