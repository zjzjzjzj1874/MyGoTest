package main

import "fmt"

const (
	one = 1 << iota
	two
)

const (
	greet = "Hello,world."
	One   = 1 << iota
	Two
)

func main() {
	//fmt.Println(one, two)
	fmt.Println(One, Two)
}
