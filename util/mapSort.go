package util

import (
	"fmt"
	"sort"
)

type People struct {
	Name string `json:"name"`
	Age  int64  `json:"age,omitempty"`
}

// 重写sort方法 -- 现在可以对person的数组按照age字段进行排序了
type PersonSlice []People

func (s PersonSlice) Len() int { return len(s) }
func (s PersonSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Age > s[j].Age }

func SortPersonMap() {
	a := PersonSlice{
		{
			Name: "AAA",
			Age:  5,
		},
		{
			Name: "BBB",
			Age:  32,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	sort.Stable(a)
	fmt.Println(a)
}
