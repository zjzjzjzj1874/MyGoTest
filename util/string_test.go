package util

import (
	"fmt"
	"testing"
)

func TestIsChineseChar(t *testing.T) {
	fmt.Println(IsChineseChar("ajesgoiab"))
	fmt.Println(IsChineseChar("asj指标考核"))
	fmt.Println(IsChineseChar("把客户ajesgoiab"))
}
