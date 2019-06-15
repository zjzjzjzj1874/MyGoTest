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
