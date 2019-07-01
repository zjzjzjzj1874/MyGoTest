package log

import (
	"os"
	"time"
)

// 写log
func WriteLog(text string) {
	path := "errLog" + time.Now().Format("20060102") + ".log"
	if isExist(path) == false{
		create(path)
	}
	// 文件是否存在
	fMode := os.O_CREATE
	f, _ := os.OpenFile(path,  fMode, 0666)
	defer f.Close()
	f.WriteString(time.Now().Format("2006-01-02 15:04:05")+":"+text+"\n")
}

// 创建文件
func create(path string) interface{} {
	_, err := os.Create(path)
	if err != nil {
		return err
	}else{
		return true
	}
}
// 文件（夹）是否存在
func isExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}