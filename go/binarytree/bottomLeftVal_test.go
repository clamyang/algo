package binarytree

import "testing"

func TestBottomLeftVal(t *testing.T) {
	root := NewTree()
	root.init()

	BottomLeftVal(root)
}
