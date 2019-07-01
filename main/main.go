package main

import (
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func main() {

	WriteFile("测试添加")

	//CheckNumberAndLetter("abcsdlj","0aA")
	//n := time.Now()
	//timeString := n.AddDate(0,0,-1).Format("2006-01-02")
	//yesterday := timeString + " 22:00:00"
	//fmt.Println("yesterday ==== ",yesterday)
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// this function is testing the type of each char in STR.
// if the checkType contains "1",the STR can contains 1-9;
// if the checkType contains "a",the STR can contains a-z;
// if the checkType contains "A",the STR can contains A-Z;
func CheckNumberAndLetter(str, checkType string) bool{
	reg := "^["
	if strings.Contains(checkType,"1"){
		reg += "0-9"
	}
	if strings.Contains(checkType,"a"){
		reg += "a-z"
	}
	if strings.Contains(checkType,"A"){
		reg += "A-Z"
	}
	reg += "]$"
	for i := 0;i < len(str);i ++ {
		if !regexp.MustCompile(reg).MatchString(string(str[i])){
			return false
		}
	}
	return true
}



//写log
func WriteFile(text string) {
	path := "errLog" + time.Now().Format("20060102")
	if isExist(path) == false{
		Create(path)
	}
	//文件是否存在
	fMode := os.O_CREATE
	f, _ := os.OpenFile(path,  fMode, 0666)
	defer f.Close()
	f.WriteString(time.Now().Format("2006-01-02 15:04:05")+"\n"+text+"\n")
}


// 创建文件
func Create(path string) interface{} {
	_, err := os.Create(path)
	if err != nil {
		return err
	}else{
		return true
	}
}

//  文件（夹）是否存在
func isExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}