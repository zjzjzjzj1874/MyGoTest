package RregExp

import (
	"fmt"
	"testing"
)

func TestCheckNumberAndLetter(t *testing.T) {
	fmt.Println(CheckNumberAndLetter("Hello","1aA"))
	fmt.Println(CheckNumberAndLetter("Hello","a"))
	fmt.Println(CheckNumberAndLetter("Hello","1"))
	fmt.Println(CheckNumberAndLetter("Hello","A"))
	fmt.Println(CheckNumberAndLetter("alhjebfjb","a"))
	fmt.Println(CheckNumberAndLetter("1236542342","1"))
	fmt.Println(CheckNumberAndLetter("2983333aweaswdva907YUVuyb","1aA"))
}
