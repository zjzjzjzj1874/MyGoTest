package main

import (
	"fmt"
	"sort"
)

func main() {
	SortPersonMap()
	//runtime.GOMAXPROCS(1)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("i1: ", i)
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("i2: ", i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

}
func TestIss0615()  {
	fmt.Println("0615:新增分支:iss0615,测试分支上的提交")
}

type Person struct {
	Name        string `json:"name"`
	Age         int64  `json:"age,omitempty"`
}

// 重写sort方法 -- 现在可以对person的数组按照age字段进行排序了
type PersonSlice []Person

func (s PersonSlice) Len() int           { return len(s) }
func (s PersonSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Age > s[j].Age }

func SortPersonMap()  {
	a := PersonSlice {
		{
			Name: "AAA",
			Age: 5,
		},
		{
			Name: "BBB",
			Age: 32,
		},
		{
			Name: "CCC",
			Age: 0,
		},
		{
			Name: "DDD",
			Age: 22,
		},
		{
			Name: "EEE",
			Age: 11,
		},
	}
	sort.Stable(a)
	fmt.Println(a)
}