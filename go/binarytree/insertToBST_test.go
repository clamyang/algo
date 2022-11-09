package binarytree

import (
	"fmt"
	"testing"
)

func TestInsertToBSTByIter(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	InsertToBSTByIter(root, 9)

	fmt.Println(TraverseInLayer(root))
}
