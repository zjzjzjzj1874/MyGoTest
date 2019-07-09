package dataStructure

import (
	"fmt"
	"testing"
)

func TestLinkedList_IsEmpty(t *testing.T) {
	l := NewLinkedList()
	fmt.Println(l)
	fmt.Println(l.IsEmpty())
}
