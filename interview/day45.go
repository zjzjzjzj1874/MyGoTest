package main

import "fmt"

func main() {
	//varRepeat()

	//specialType()

	//arrRange()
	//pointRange()
	//sliceRange()
	sliceRan()
}

// region 变量的重复申明
func varRepeat() {
	// region 例1 => false
	//one := 1 // unused variable 'one'=>申明变量未使用
	//one := 1 // no new variables on left side of := =>
	// endregion 例1

	// region 例2 => false
	//two := 2 // unused variable 'two'=>申明变量未使用
	//two = 2
	// endregion 例2

	// region 例3 => true
	three := 3 // unused variable 'two'=>申明变量未使用
	three = 3
	fmt.Println(three)
	// endregion 例3

}

// endregion 变量的重复申明

// region 数组,对象换行与否
func lineFeed() {
	// region 例1 false
	x := []int{
		1,
		//2 // 缺少一个逗号
	}
	_ = x
	// endregion 例1

	// region 例2 一下三种是正确的 true
	y := []int{
		1,
		2,
	}
	z := []int{1, 2}
	zz := []int{1, 2}
	fmt.Println(y, z, zz)
	// endregion 例2
}

// endregion 数组,对象换行与否

// region 特殊类型别名
func printIt(x byte) {
	fmt.Println(x)
}
func specialType() {
	var a byte = 0x11 // 16进制的数
	var b uint8 = a
	var c uint8 = a + b
	printIt(a)
	printIt(c)
	// 注意:与 rune 是 int32 的别名一样，byte 是 uint8 的别名，
	// 别名类型无序转换，可直接转换。
}

// endregion 特殊类型别名

// region 加餐:range的使用

type T struct {
	n int
}

// 数组的遍历,是值拷贝的遍历;
func arrRange() {
	ts := [2]T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3     // 这里改变的其实是拷贝的副本的值,原数组中的值不变
			ts[1].n = 9 // 这是通过角标访问原数组,并改变原数组
		case 1:
			fmt.Println(t.n, " ") // 这里并不是9,因为在循环时,这个副本的值都已经确定了
		}
	}
	fmt.Println(ts) // 打印 => 0 {0 9}
}

// 数组指针遍历
func pointRange() {
	ts := [2]T{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3 // t其实是遍历中的一个副本，修改它的值不会影响原数组的值
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts) // 打印 => 9 {0 9}
}

// 切片的遍历
func sliceRange() {
	ts := [2]T{}
	for i := range ts[:] {
		t := &ts[i] // 这步最关键:t其实是切片的一个元素的地址,所以和原元素共享数据,修改t也会修改原数据
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts) // 打印 => 9 {3 9}
}

// 直接生成切片
func sliceRan() {
	//ts := [2]T{}
	ts := make([]T, 2)
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts) // 打印 => 9 {0 9}
}

// ------------- 总结 ------------
// 当使用for-range遍历一个容器时,其实遍历的是容器的一个副本
// 1.遍历数组时,则会复制这个数组的所有元素;==>这个时候注意操作的是原数组还是复制的数组;
// 2.复制一个切片(或数组的指针),此切片(或数组)不会被复制,则此副本切片(或数组)将和原切片(或数组)共享元素;
// 3.分析这类问题时,主要是确认到底用的是原数组还是副本,地址还是值;

// endregion 加餐:range的使用
