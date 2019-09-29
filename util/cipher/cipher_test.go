package cipher

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncryptWithAES(t *testing.T) {
	fmt.Println(EncryptWithAES("953"))
}

func TestDecodeWithAES(t *testing.T) {
	fmt.Println(DecodeWithAES("84512f"))
}

func TestMd5Encode(t *testing.T) {
	fmt.Println("原始密文:abc,加密后:", Md5Encode("abc"))
	fmt.Println(GetMd5String("abc"))
}

func TestGetGuid(t *testing.T) {
	fmt.Println("", GetGuid())
}

func TestBase64Encode(t *testing.T) {
	fmt.Println("原文:953,Base64加密:", Base64Encode("953"))
	fmt.Println(Base64Decode(Base64Encode("abc")))
}

func TestBase64Decode(t *testing.T) {
	fmt.Println(Base64Decode("emgub3VqLmlhbjk1MzA5Lk8yNS5PMjAuNzg="))
}

func TestRsaDecrypt(t *testing.T) {
	//s ,err := RsaEncrypt([]byte("953"))
	//fmt.Println(s,err)
	fmt.Println("密文长度 ==>", len("EFWEZcJoPrmWy0bHXhP0CozVZe9099RhDz8PWJyvbvGZ5Q2uXIF9ArGYZDT32rMEWZEG4NXsgSWUAG7M30gqpo/tWSl+gPzSF+qrBPWTkO9xTHk5lbcEQSNFYIYFWQZF0rscyXWnaAmEYJ3Psf/MgRZlAYMrwAe3r+fhEbn+qGY="))

	//生成密钥文件
	//fmt.Println("start")
	//GenRsaKey(1024)
	//fmt.Println("end")
	//加密解密测试
	data, _ := RsaEncrypt([]byte("953"))
	fmt.Println("加密 ==>", string(data))
	data2 := string(data)
	fmt.Println("密文 ==>", base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt([]byte(data2))
	fmt.Println("解密 ==>", string(origData))
}
