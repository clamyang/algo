package binarytree

import (
	"testing"
)

func TestBinaryTreePath(t *testing.T) {
	root := NewTree()
	root.initComplete()

	paths := make([]int, 0, 100)
	TreePaths(root, paths)
}
