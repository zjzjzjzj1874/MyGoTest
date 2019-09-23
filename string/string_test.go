package string

import (
	"fmt"
	"testing"
)

//测试Reverse函数是否生效
func TestReverse(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"back", "kcab"},
		{"hello", "olleh"},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q,want == %q", c.s, got, c.want)
		}
	}
}
func TestReverse2(t *testing.T) {
	fmt.Println(reverse1("hello"))
	fmt.Println(reverse2("hello"))
	fmt.Println(Reverse("hello"))
}
