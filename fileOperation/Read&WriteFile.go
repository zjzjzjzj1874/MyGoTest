package fileOperation

import (
	"fmt"
	"io/ioutil"
	"os"
)

//创建新的文件
func CreateFile(fileName, fileType string) {
	newFile, err := os.Create(fileName + "." + fileType)
	if err != nil {
		fmt.Println("\r\n fail to create a file,err === ", err)
	}
	fmt.Println("new file == ", newFile)
	newFile.Close()
}

//写入文件
func WriteFile(fileName, fileContent string) (err error) {
	data := []byte(fileContent)
	return ioutil.WriteFile(fileName, data, 0644)
}

//检查文件是否存在
func IsExist(fileName, fileType string) {
	file, OpenErr := os.Open(fileName + "." + fileType) //这个文件是在main包下面
	if OpenErr != nil {
		fmt.Println("\r\n open file occurred an err === ", OpenErr)
	}
	defer file.Close()
	fileInfo, InfoErr := file.Stat()
	if InfoErr != nil {
		fmt.Println("\r\n Info has an err ==== ", InfoErr)
	}
	fmt.Println("\r\n fileInfo === ", fileInfo)
}
