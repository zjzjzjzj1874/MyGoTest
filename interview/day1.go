package interview

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
