package binarytree

import (
	"fmt"
	"testing"
)

func TestConvertBSTToGreaterTree(t *testing.T) {
	root := NewTree()
	root.initSearchTree()
	
	ConvertBSTToGreaterTree(root)
	fmt.Println(TraverseInLayer(root))
}
