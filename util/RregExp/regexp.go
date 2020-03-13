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
func CheckNumberAndLetter(str, checkType string) bool {
	reg := "^["
	if strings.Contains(checkType, "1") {
		reg += "0-9"
	}
	if strings.Contains(checkType, "a") {
		reg += "a-z"
	}
	if strings.Contains(checkType, "A") {
		reg += "A-Z"
	}
	reg += "]$"
	for i := 0; i < len(str); i++ {
		if !regexp.MustCompile(reg).MatchString(string(str[i])) {
			return false
		}
	}
	return true
}

// 正则匹配电话 简单匹配
func IsPhoneNumber(phone string) bool {
	return regexp.MustCompile(`^1((45|47)|[356789]\d)\d{8}$`).MatchString(phone)
}

// 正则匹配密码
func IsCorrectPassword(password string) bool {
	PasswordRex := regexp.MustCompile(`^[\w+=-]{6,20}$`)
	return PasswordRex.MatchString(password)
}

// 正则匹配身份证 (二代身份证验证)
func IsCorrectIdCard(idCard string) bool {
	PasswordRex := regexp.MustCompile(`^[1-9]\d{5}(((1[89]|20)\d{2}(((0[13578]|1[0-2])(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)(0[1-9]|[12][0-9]|30))|(02(0[1-9]|[1][0-9]|2[0-8]))))|((((1[89]|20)(0[48]|[2468][048]|[13579][26]))|((19|20)00))0229))\d{3}(\d|X|x)$`)
	return PasswordRex.MatchString(idCard)
}

// 银行卡号简单匹配 注:严格匹配可以用支付宝接口
// 我国银行卡片主要分两种——借记卡（普通储蓄卡）和信用卡（贷记卡）。借记卡的卡号一般是19位，但有例外：招行16位，交行17位；信用卡的卡号一般是16位。
func IsCorrectBankNo(bankNo string) bool {
	bankNoRex := regexp.MustCompile(`^([1-9])(\d{15}|\d{16}|\d{18})$`)
	return bankNoRex.MatchString(bankNo)
}
