package dataStructure

import "testing"

func TestDoubleLinkedList_Append(t *testing.T) {
	d := NewDoubleLinkedList()
	d.Append(1)
	d.Append(3)
	d.Traverse()
}
