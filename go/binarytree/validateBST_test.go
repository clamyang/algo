package binarytree

import (
	"fmt"
	"sort"
	"testing"
)

func TestValidateBSTByIn(t *testing.T) {
	root := NewTree()
	root.initSearchTree()

	ValidateBSTByIn(root)
	fmt.Println(sort.IsSorted(arrIn))
}
