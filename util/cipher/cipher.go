package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	//"golang.org/x/crypto/scrypt"
	"io"
	"strings"
)

// md5加密
func Md5Encode(data string) string {
	raw := []byte(data)
	md5Ctx := md5.New()
	md5Ctx.Write(raw)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

/*
 * AES加密解密
 */
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// 创建加密算法
func createAES() cipher.Block {
	// AES密匙：也可以将这个密匙当做盐存到数据库里面
	var privateKey = "thisFunctionIsMadeByZJZJ" //16个字符
	c, err := aes.NewCipher([]byte(privateKey))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(privateKey), err)
		return nil
	}
	return c
}

// 使用AES加密
func EncryptWithAES(str string) string {
	plaintext := []byte(str)
	c := createAES()
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)
	return fmt.Sprintf("%x", cipherText)
}

// 使用AES解密
func DecodeWithAES(str string) string {
	cipherText, _ := hex.DecodeString(str)
	c := createAES()
	cfbDec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(cipherText))
	cfbDec.XORKeyStream(plaintextCopy, cipherText)
	return string(plaintextCopy)
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

// 加密
func Base64Encode(str string) string {
	var src = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

// 解密
func Base64Decode(str string) (string, error) {
	var src = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`  
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCpf7pZbKxCKuxqYdoanknSDvD8tJp7FASPZymFMFrJsEDHc9j5
LBIxr9rd70UWApyOtrQgCeMFw6bfD6fF/0WMtzveLQ3qG2bUwYQm9Lr2RNL7nydy
qhVg3HC9zulTiMfkERZfvjZdhbVo9NXwDKVsFWrrk/5vQiJ8EjPd63/RlwIDAQAB
AoGAE1XSuCjBbbrfxTLsYmT0HtY9f1ZK2QdrjcBC6EKf2KoWeaopciMo4CojWXXV
97DMkyscWRtHnny3KHLsvJVmJXuotDn3l3xJje4GTcaiGy7f5MH42Aq2hNUjiraU
lvrNiYYq1e6RL4VS0YLc+C4XEgsGE/lBrwom+mQSkP3KMiECQQDMNv3HaBJn/h26
wRVzFnMl0H4Y+RLiqT+fNKHqFquf9u/RWvtof2MV+B28yNIo3ufrIdT5MMEFl6H4
31tZ1BxrAkEA1HsWRb1aGhonfpqrKBE0mlDZx3Y3HRQ1qOG2dk+FRsLVtuKP+xHb
Hvun5CO/9ZuAq08DLXF9Ebz3CFAXK8qqhQJBAMZY6yjQ9n+3G90WSOUdev3RgYhz
81nflYHmtxUMq+mVCN1JB0M5512hPhDs5OL5jjydAaR/LBtoadO17Z5UHL0CQDwL
bIfYspWdvntwid2QvyS8pE5RgdGd3GwVHNLiNe+BL5O3AqkYqqtewlseHyjxALNo
aKV25LkWhVi8CVA+vWECQGRNKk+sHaljL+/aaxx4sBrBFcq70RQLhW/2mkeYYTom
ztHEIWP5EVSrnzdAMecov3tXtvC4WZqMlwbeHB2L7uQ=
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`  
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpf7pZbKxCKuxqYdoanknSDvD8
tJp7FASPZymFMFrJsEDHc9j5LBIxr9rd70UWApyOtrQgCeMFw6bfD6fF/0WMtzve
LQ3qG2bUwYQm9Lr2RNL7nydyqhVg3HC9zulTiMfkERZfvjZdhbVo9NXwDKVsFWrr
k/5vQiJ8EjPd63/RlwIDAQAB
-----END PUBLIC KEY-----    
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
