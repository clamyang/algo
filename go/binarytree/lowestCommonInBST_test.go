package binarytree

import (
	"fmt"
	"testing"
)

func TestLowestCommonInBST(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	fmt.Println(LowestCommonInBST(root, root.left.right, root.left.left))
}

func TestLowestCommonInBSTByIter(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	fmt.Println(LowestCommonInBSTByIter(root, root.left.right, root.left.left))
}

