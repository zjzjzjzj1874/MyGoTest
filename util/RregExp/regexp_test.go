package RregExp

import (
	"fmt"
	"testing"
)

func TestCheckNumberAndLetter(t *testing.T) {
	fmt.Println(CheckNumberAndLetter("Hello", "1aA"))
	fmt.Println(CheckNumberAndLetter("Hello", "a"))
	fmt.Println(CheckNumberAndLetter("Hello", "1"))
	fmt.Println(CheckNumberAndLetter("Hello", "A"))
	fmt.Println(CheckNumberAndLetter("alhjebfjb", "a"))
	fmt.Println(CheckNumberAndLetter("1236542342", "1"))
	fmt.Println(CheckNumberAndLetter("2983333aweaswdva907YUVuyb", "1aA"))
}

func TestIsPhoneNumber(t *testing.T) {
	fmt.Println(IsPhoneNumber("1523"))
	fmt.Println(IsPhoneNumber("1552548721"))
	fmt.Println(IsPhoneNumber("14525487212"))
	fmt.Println(IsPhoneNumber("14725487212"))
	fmt.Println(IsPhoneNumber("14325487212"))
	fmt.Println(IsPhoneNumber("15525487212"))
	fmt.Println(IsPhoneNumber("18725487212"))
}

func TestIsCorrectPassword(t *testing.T) {
	fmt.Println(IsCorrectPassword(""))
	fmt.Println(IsCorrectPassword("1as2"))
	fmt.Println(IsCorrectPassword("1as22341"))
	fmt.Println(IsCorrectPassword("Aas22341"))
	fmt.Println(IsCorrectPassword("Aas221"))
	fmt.Println(IsCorrectPassword("Aas周21"))
	fmt.Println(IsCorrectPassword("Aas22"))
}

func TestIsCorrectIdCrad(t *testing.T) {
	fmt.Println(IsCorrectIdCard("036767199903037878")) // 0开头
	fmt.Println(IsCorrectIdCard("130981199312253466"))
	fmt.Println(IsCorrectIdCard("130981200002293466")) // 2月29日出生
	fmt.Println(IsCorrectIdCard("130981200002303466")) // 2月30日出生
	fmt.Println(IsCorrectIdCard("13098119931225346x"))
	fmt.Println(IsCorrectIdCard("13098119931225346")) // 17位
}

func TestIsCorrectBankNo(t *testing.T) {
	fmt.Println(IsCorrectBankNo("1234567890123456789")) // 19位
	fmt.Println(IsCorrectBankNo("123456789012345678"))  // 18位
	fmt.Println(IsCorrectBankNo("12345678901234567"))   // 17位
	fmt.Println(IsCorrectBankNo("1234567890123456"))    // 16位
	fmt.Println(IsCorrectBankNo("123456789012345"))     // 15位
	fmt.Println(IsCorrectBankNo("0309811993122534"))    // 0开头16位
}
