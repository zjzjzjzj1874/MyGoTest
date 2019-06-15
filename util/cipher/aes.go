package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func main() {
	d := []byte("ABC")
	key := []byte("zhoujian19940104")
	fmt.Println("pwd:", string(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\r\n x1 === ", string(x1))
	fmt.Println("before encipher:", x1)
	x2, err := decryptAES(x1, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("after deciphering:", string(x2))
}
