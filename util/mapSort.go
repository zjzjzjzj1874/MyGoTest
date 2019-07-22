package util

import (
	"fmt"
	"sort"
)

type People struct {
	Name        string `json:"name"`          // json:"name" 表示在json中的key == name;如果没有的话,json中的key = Name
	Age         int64  `json:"age,omitempty"` // omitempty 忽略空值:如果没有该关键字,key为空的时候也会显示出来;有的话,key为空不会出现该字段的相关信息
	Tel         string `json:"tel,omitempty"`
	Pwd1        string `json:"-"`       // json:"-" 表示忽略这个字段;(对于密码这种字段起到保护的作用,不输出)
	Pwd2        string `json:"-,"`      // json:"-," 表示不会忽略这个字段;并且key显示为 "-"
	Int64String int64  `json:",string"` // json:",string" 将该类型转化为string;
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
