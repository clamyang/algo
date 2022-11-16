package binarytree

import (
	"fmt"
	"testing"
)

func TestInsertIntoBSTByIter(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	InsertIntoBSTByIter(root, 9)

	fmt.Println(TraverseInLayer(root))
}

func TestInsertIntoBSTByRecursive(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	InsertIntoBSTByRecursive(root, 9)
	fmt.Println(TraverseInLayer(root))
}
