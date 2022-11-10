package binarytree

import (
	"fmt"
	"testing"
)

func TestDeleteNodeInBST(t *testing.T) {
	root := NewTree()
	root.initSearchTree()


	fmt.Println(TraverseInLayer(DeleteNodeInBST(root, 4)))
}
