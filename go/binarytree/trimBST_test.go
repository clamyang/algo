package binarytree

import (
	"fmt"
	"testing"
)

func TestTrimBST(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	fmt.Println(TraverseInLayer(TrimBST(root, 6, 8)))
}
