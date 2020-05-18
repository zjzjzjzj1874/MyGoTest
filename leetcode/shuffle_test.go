package leetcode

import (
	"fmt"
	"testing"
)

func TestShuffleTime(t *testing.T) {
	var cnt1 = map[int]int{}
	//for i := 0; i < 10000000; i++ {
	//	c1 := shuffle1()
	//	cnt1[int(c1[0])]++
	//}

	var cnt2 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		c2 := shuffle2()
		cnt2[int(c2[0])]++
	}

	fmt.Println(cnt1, "\n", cnt2)
}

func TestTransfer(t *testing.T) {
	transfer()
}
