package util

import (
	"container/list"
	"fmt"
	"testing"
)

func TestConvertByteSliceToIntSlice(t *testing.T) {
	//byteSlice := []byte{0x01,0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,0x31,}
	//in := ConvertByteSliceToIntSlice(byteSlice)
	//fmt.Println(in)
	//fmt.Println(ConvertIntSliceToByteSlice(in))
	//fmt.Println(ConvertIntSliceToByteSlice(in))
	//
	//fmt.Println(TestBit(2,2))

	//v := 1
	//fmt.Println(1&2)
	//fmt.Println(100&2)
	//fmt.Println(200&2)

	l := list.New()
	e1 := l.PushFront(123)
	l.InsertAfter(12345, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//for i := 0; i < 10; i++ {
	//	x := v | i
	//	fmt.Printf("i = %d => x = 1|%d == %d,v = %d,\n", i, i, x, v)
	//}

	//b := byte(4)
	//fmt.Println(int(b))
	//
	//if b > 3{
	//	fmt.Println("我大于3哦")
	//
	//	goto
	//	fmt.Println("我大于3哦")
	//}

	//max := 1
	//min := 4
	//fmt.Println(max|min)
	//
	//fmt.Println(3/16)
	//fmt.Println(13/16)
	//fmt.Println(23/16)
	//fmt.Println(33/16)
	//fmt.Println(43/16)
	//fmt.Println(53/16)

	//m := map[byte]struct{}{}
	//m[2] = struct{}{}
	//fmt.Println(m,len(m))
	//m[3] = struct{}{}
	//fmt.Println(m,len(m))

}
