package log

import (
	"MyGoTest/util"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// 写错误log
func ErrorLog(i... interface{}) {
	path := "./log/errLog" + time.Now().Format("20060102") + ".log"
	if !isExist(path){
		create(path)
	}
	fMode := os.O_APPEND|os.O_WRONLY
	f, _ := os.OpenFile(path,fMode,0666)
	defer f.Close()
	_,err := f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [ERROR] | " + util.InterfaceToString(i) + "\n")
	fmt.Println(err)
}
// 写普通打印log
func InfoLog(i... interface{}) {
	path := "./log/infoLog" + time.Now().Format("20060102") + ".log"
	if !isExist(path){
		create(path)
	}
	fMode := os.O_APPEND|os.O_WRONLY
	f, _ := os.OpenFile(path,fMode,0666)
	defer f.Close()
	_,err := f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [INFO] | " + util.InterfaceToString(i) + "\n")
	fmt.Println(err)
}
//将可能是int\int32\int64\float64\string类型的interface{}转换成string，可继续完善类型
func interfaceToString(i interface{}) string {
	switch i.(type) {
	case int:
		return strconv.Itoa(i.(int))
	case int32:
		return strconv.Itoa(int(i.(int32)))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case string:
		return i.(string)
	default:
		str,_ := json.Marshal(i)
		return string(str)
	}
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