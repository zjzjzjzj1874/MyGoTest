package interview

import "fmt"

func main() {
	abc := []int{1, 2, 3, 4, 5, 6}
	var x int = 3
	var i int = 100
	add(abc, x, i)

}

func add(a []int, b int, c int) (d []int) {
	arr1 := new([]int)

	*arr1 = append(*arr1, a[:b]...)

	*arr1 = append(*arr1, c)

	*arr1 = append(*arr1, a[b:]...)

	fmt.Println(a)
	fmt.Println(*arr1)
	return a
}
func add1(a []int, b int, c int) (d []int) {
	abcdef := a
	var abcd []int
	abcd = append(abcdef[:b], c)
	fmt.Println(abcd)
	fmt.Println(a)
	abcd = append(abcd, a[b:]...)
	return a
}
