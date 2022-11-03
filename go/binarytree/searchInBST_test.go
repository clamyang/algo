package binarytree

import (
	"fmt"
	"testing"
)

func TestSearchInBST(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	subRoot := SearchInBST(root, 1)
	if subRoot != nil {
		fmt.Println(subRoot.val)
	} else {
		fmt.Println(3333333)
	}
}
