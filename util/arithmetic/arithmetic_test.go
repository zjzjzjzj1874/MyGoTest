package arithmetic

import (
	"fmt"
	"testing"
)

// arr := []int{1,2,3,4,5,6,10,13,14,15,16,20,23,26,28,29,30,32,33,36,38,40,41,44,47,48,50,52,55,57,58,60,61,65,234,754,1234,2432,2654,2986}
// 测试二分法查找数
func TestGetIndexByDichotomy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 10, 13, 14, 15, 16, 20, 23, 26, 28, 29, 30, 32, 33, 36, 38, 40, 41, 44, 47, 48, 50, 52, 55, 57, 58, 60, 61, 65, 234, 754, 1234, 2432, 2654, 2986}
	index := GetIndexByDichotomy(arr, 38)
	fmt.Println("index === ", index)
	fmt.Println("aim === ", arr[index])
	fmt.Println("length === ", len(arr))
}

// 冒泡排序测试
func TestBubbleSort(t *testing.T) {
	arr := []int{754, 1234, 1, 2, 55, 57, 47, 3, 30, 41, 44, 2432, 2654, 50, 52, 32, 4, 5, 23, 26, 13, 14, 15, 16, 20, 29, 33, 36, 38, 40, 48, 58, 60, 61, 28, 6, 10, 65, 234, 2986}
	fmt.Println(BubbleSort(arr, ""))
}

// 选择问题测试
func TestSelectionProblem(t *testing.T) {
	arr := []int{754, 1234, 1, 2, 55, 57, 47, 3, 30, 41, 44, 2432, 2654, 50, 52, 32, 4, 5, 23, 26, 13, 14, 15, 16, 20, 29, 33, 36, 38, 40, 48, 58, 60, 61, 28, 6, 10, 65, 234, 2986}
	fmt.Println(SelectionProblem(5, arr))
}

// 测试斐波那契数列求和
func TestFactorial(t *testing.T) {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println("this i == ", Fibonaci(i))
	}

	fmt.Println(Fibonaci(10))
}
