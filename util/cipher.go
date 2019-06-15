package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	//"golang.org/x/crypto/scrypt"
	"io"
	"strings"
)

//md5加密
func Md5Encode(data string) string {

	raw := []byte(data)
	md5Ctx := md5.New()
	md5Ctx.Write(raw)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

/*
 * Scrypt加密,需要Salt
 * 逻辑:将salt和通过salt加密后的密码持久化;以后用户登录或者其他操作的时候再用salt对post过来的密码加密,并和持久化的进行对比;
 * 前台最好也加密再传过来
 */
//非对称加密
//func EncryptByScrypt(str, salt string) string {
//	dk, err := scrypt.Key([]byte(str), []byte(salt), 32768, 8, 1, 18)
//	if err != nil {
//		fmt.Println("Scrypt加密 err:+++++++++", err)
//		return ""
//	}
//	return base64.URLEncoding.EncodeToString(dk)
//}

/*
 * AES加密解密
 */
//创建加密算法
func createAES() cipher.Block {
	//AES密匙
	//var key_text = "zhouzhoujianjianzhouzhoujianjian" //32个字符
	//var key_text = "zhouzhouzhoujianjianjian" //24个字符
	var key_text = "zhoujianjianzhou" //16个字符
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		return nil
	}
	return c
}

//使用AES加密
func EncryptWithAES(str string) string {
	plaintext := []byte(str)
	c := createAES()
	var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)
	return fmt.Sprintf("%x", cipherText)
}

//使用AES解密
func DecodeWithAES(str string) string {
	cipherText, _ := hex.DecodeString(str)
	c := createAES()
	var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	cfbDec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(cipherText))
	cfbDec.XORKeyStream(plaintextCopy, cipherText)
	return string(plaintextCopy)
}

/**
 * 获取一个Guid值
 */
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s + "zhoujian")) //使用zhoujian名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * base64的加密与解密
 */
const (
	//BASE64字符表,不要有重复
	base64Table        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_" //需要64个不重复的字符
	hashFunctionHeader = "zh.ouj.ian"
	hashFunctionFooter = "09.O25.O20.78"
)

var coder = base64.NewEncoding(base64Table)

/**
 * base64加密
 */
func Base64Encode(str string) string {
	var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

/**
 * base64解密
 */
func Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}
