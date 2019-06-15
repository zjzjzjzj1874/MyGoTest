package fileOperation

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//检查文件是否存在
func FileExist(fileName string) bool {
	_, err := os.Lstat(fileName)
	return !os.IsNotExist(err)
}

//整个文件读取
//按字节读取，将整个文件读取到缓冲区buffer
func ReadWholeFile(fileName, fileType string) {
	file, OpenErr := os.Open(fileName + "." + fileType) //这个文件是在main包下面
	if OpenErr != nil {
		fmt.Println("\r\n open file occurred an err === ", OpenErr)
	}
	defer file.Close()
	fileInfo, InfoErr := file.Stat()
	if InfoErr != nil {
		fmt.Println("\r\n Info has an err ==== ", InfoErr)
	}
	fmt.Println("\r\n  fileInfo ==== ", fileInfo)
	fmt.Println("\r\n  file name ==== ", fileInfo.Name())
	fmt.Println("\r\n  是否是文件夹 ==== ", fileInfo.IsDir())
	fmt.Println("\r\n  file Mode ==== ", fileInfo.Mode())
	fmt.Println("\r\n  file size ==== ", fileInfo.Size())
	fmt.Println("\r\n  file sys ==== ", fileInfo.Sys())
	fmt.Println("\r\n  file ModTime ==== ", fileInfo.ModTime())

	buffer := make([]byte, fileInfo.Size()) //缓冲区,长度和文件的长度一致
	byteString, ReadErr := file.Read(buffer)
	if ReadErr != nil {
		fmt.Println("read err ===== ", ReadErr)
	}
	fmt.Println("bytes read === ", byteString)
	fmt.Println("byteStream to string  ======================  ", string(buffer))
}

//按规定的长度读取
func ReadFileBySpecialLength(fileName, fileType string, length int64) {
	file, OpenErr := os.Open(fileName + "." + fileType) //这个文件是在main包下面
	if OpenErr != nil {
		fmt.Println("\r\n open file occurred an err === ", OpenErr)
	}
	defer file.Close()

	buffer := make([]byte, length) //缓冲区,长度和文件的长度一致
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		fmt.Println(string(buffer[:bytesRead]))
	}
}

//逐行读取
func ReadFileByRow(fileName, fileType string) {
	file, OpenErr := os.Open(fileName + "." + fileType) //这个文件是在main包下面
	if OpenErr != nil {
		fmt.Println("\r\n open file occurred an err === ", OpenErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	/*
		ScanLines (默认)
		ScanWords
		ScanRunes (遍历UTF-8字符非常有用)
		ScanBytes
	*/
	//是否有下一行
	for scanner.Scan() {
		fmt.Println("read string:", scanner.Text())
	}
}

//ioUtil读取文件
func ReadFileByIoUtil(fileName, fileType string) {
	bytes, err := ioutil.ReadFile(fileName + "." + fileType)
	if err != nil {
		fmt.Println("read file occurred an err: ", err)
	}
	fmt.Println("total bytes read：", len(bytes))
	fmt.Println("string read:", string(bytes))
}

//读取文件夹;
func ReadFileDir(dirName string) {
	fileInfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("\r\n read dir err :=============== ", err)
	}
	for _, v := range fileInfo {
		fmt.Println("file name === ", v.Name())
		fmt.Println("file size === ", v.Size())
	}
}
