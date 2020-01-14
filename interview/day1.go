package main

import "fmt"

func iscanpass1() {
	list := new([]int)

	*list = append(*list, 1) // 因为new出来的就是指针类型
	//list = append(list,1)	// 不通过
	fmt.Println(list)
}

func iscanpass2() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5}

	// arr1 = append(arr1, arr2)	// 不通过,append是对元素进行增加
	arr1 = append(arr1, arr2...)
	fmt.Println(arr1)
}

var (
	//max := 24 // 有了 var不能使用 ':=';全局变量也不能使用 :=
	//max_size := max * 2	//
	max      = 24
	max_size = max * 2
)

func iscanpass3() {
	fmt.Println(max, max_size)
}

// 对slice进行操作
func sliceOperate() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	copyOne := make([]int, len(arr)-1) // 注意：这里使用copy函数，其len最好和原数组一样大；否则:超过len的,后面的元素不会被copy
	copy(copyOne, arr)

	//copyOne = copyOne[1:]
	fmt.Println(copyOne)
}

func main() {
	sliceOperate()
}
