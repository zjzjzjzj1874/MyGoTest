package RregExp

import (
	"regexp"
	"strings"
)

// 正则验证str中是否只有数字,大小写字母
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
