package main

import (
	"fmt"
	"io"
	"os"
)

var he int

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	copyOne := make([]int, len(arr))
	copy(copyOne, arr)

	copyOne = copyOne[1:]
	fmt.Println(copyOne)

	// 这里创建的一个变量he,返回的是1,但是出来以后,这个变量已经是2了;
	//fmt.Println(deferFunc())
	//fmt.Println(he)

	// 写入文件
	//_ = ioutil.WriteFile("test.txt", []byte("Hello world."), 0644)
	//con, err := Contents("test.txt")
	//fmt.Println(con, err)

	// 判断变量类型
	//var t interface{}
	//t = "hello"
	////t = functionOfSomeType()
	//switch t := t.(type) {
	//default:
	//	fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	//case bool:
	//	fmt.Printf("boolean %t\n", t) // t has type bool
	//case int:
	//	fmt.Printf("integer %d\n", t) // t has type int
	//case *bool:
	//	fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	//case *int:
	//	fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	//}
	//
	//
	//// 遍历字符串
	//for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
	//	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	//}
}

// 测试defer和return并存
func deferFunc() int {
	he = 1

	defer func() {
		he++
	}()

	return he
}

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

// 读取文件：用一个for循环来保证一定可以读取到该文件
func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

func reverse(a []int) {

	//Reverse a
	//a := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("before reverse :", a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("after  reverse :", a)

}

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
