package arithmetic

import (
	"fmt"
	"sort"
)

func GetIndexByDichotomy(arr []int, target int) int {
	return getIndexByDichotomy(arr, target)
}

// 适用范围:对于有序数组
// 二分法查找某个数 -- 排序数组
func getIndexByDichotomy(arr []int, aim int) (aimIndex int) {
	// start和end最大最小的角标
	start := 0
	end := len(arr) - 1
	for start <= end {
		aimIndex = (start + end) / 2
		if aim == arr[aimIndex] {
			return
		} else if aim > arr[aimIndex] {
			start = aimIndex + 1
		} else {
			end = aimIndex - 1
		}
	}
	return -1
}

// 选择问题:
func SelectionProblem(k int, arr []int) []int {
	return selectionProblem(k, arr)
}

// 找出数组arr中k个最大值,其中,k < len(arr);
func selectionProblem(k int, arr []int) []int {
	if k >= len(arr) {
		return arr
	}
	arr = BubbleSort(arr)
	return arr[:k]
}


//go实现斐波那契数列
func Fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonaci(n-2) + Fibonaci(n-1)
}

//递归调用函数求阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

// 两数之和
//func SumOfTwoNumber(nums []int,target int) []int {
//
//}