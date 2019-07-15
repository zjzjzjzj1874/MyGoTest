package arithmetic

import (
	"MyGoTest/util"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	var arr []int
	for i:= 10000;i < 20000;i ++{
		t,_ := strconv.Atoi(util.GetRandomString(5,"0"))
		arr = append(arr, i + t)
	}
	// 开始时间
	start := time.Now()
	fmt.Println(SelectionSort(arr))
	// 结束时间
	end := time.Now()
	//执行时间
	latency := end.Sub(start)
	fmt.Println(latency)
}

func TestBubbleSort(t *testing.T) {
	var arr []int
	for i:= 10000;i < 20000;i ++{
		t,_ := strconv.Atoi(util.GetRandomString(5,"0"))
		arr = append(arr, i + t)
	}
	start := time.Now()
	fmt.Println(BubbleSort(arr))
	// 结束时间
	end := time.Now()
	//执行时间
	latency := end.Sub(start)
	fmt.Println(latency)

}