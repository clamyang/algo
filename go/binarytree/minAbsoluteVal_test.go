package binarytree

import "testing"

func TestMinAbsoluteValByLayer(t *testing.T) {
	root := NewTree()
	root.initSearchTree()
	println(MinAbsoluteValByLayer(root))
}
