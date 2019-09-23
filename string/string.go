package string

// 字符串反转
// 1.收尾调换位置
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// 2.倒序循环输出
func reverse1(str string) (str1 string) {
	for i := len(str);i > 0; i--{
		str1 += string(str[i-1])
	}
	return
}
// 3.循环拼接字符串
func reverse2(str string) (str1 string) {
	for _, v := range str { //_ 占位使用
		str1 = string(v) + str1
	}
	return
}