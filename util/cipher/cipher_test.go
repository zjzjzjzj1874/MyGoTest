package cipher

import (
	"fmt"
	"testing"
)

func TestEncryptWithAES(t *testing.T) {
	fmt.Println(EncryptWithAES("Chinese"))
}

func TestDecodeWithAES(t *testing.T) {
	fmt.Println(DecodeWithAES("fb33441f2693c0"))
}

func TestMd5Encode(t *testing.T) {
	fmt.Println("原始密文:abc,加密后:",Md5Encode("abc"))
	fmt.Println(GetMd5String("abc"))
}

func TestGetGuid(t *testing.T) {
	fmt.Println("",GetGuid())
}

func TestBase64Encode(t *testing.T) {
	fmt.Println("原文:abc,Base64加密:",Base64Encode("abc"))
	fmt.Println(Base64Decode(Base64Encode("abc")))
}