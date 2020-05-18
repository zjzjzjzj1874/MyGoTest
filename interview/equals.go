package main

import "fmt"

// 切片和map不可以用 == 比较
func CompareSliceMapArrayChan() {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	//fmt.Println(s1 == s2) // 编译不通过: operator == not defined on []int
	fmt.Println(s1 == nil, s2 == nil)

	m1 := make(map[string]int)
	m2 := make(map[string]int)
	//fmt.Println(m1 == m2) // 编译不通过: operator == not defined on map...
	fmt.Println(m1 == nil, m2 == nil)

	a1 := [4]int{1, 3}
	a2 := [4]int{1, 2}
	fmt.Println(a1 == a2)

	c1 := make(chan int, 4)
	c2 := make(chan int, 4)
	fmt.Println(c1 == c2)

}
