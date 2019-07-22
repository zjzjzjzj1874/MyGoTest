// 简单的日志打印方法
package util

import (
	"fmt"
	"os"
	"time"
)

// 错误日志记录
func ErrorLog(i... interface{}) {
	path := "./log/errLog" + time.Now().Format("20060102") + ".log"
	if !isExist(path){
		create(path)
	}
	fMode := os.O_APPEND|os.O_WRONLY
	f, _ := os.OpenFile(path,fMode,0666)
	defer f.Close()
	_,err := f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [ERROR] | " + InterfaceToString(i) + "\n")
	fmt.Println(err)
}
// 普通日志记录
func InfoLog(i... interface{}) {
	path := "./log/infoLog" + time.Now().Format("20060102") + ".log"
	if !isExist(path){
		create(path)
	}
	fMode := os.O_APPEND|os.O_WRONLY
	f, _ := os.OpenFile(path,fMode,0666)
	defer f.Close()
	_,err := f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [INFO] | " + InterfaceToString(i) + "\n")
	fmt.Println(err)
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