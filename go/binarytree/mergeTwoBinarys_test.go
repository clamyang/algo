package binarytree

import (
	"fmt"
	"testing"
)

func TestMergeTwoBinarys(t *testing.T) {
	left := NewTree()
	left.init()
	
	right := NewTree()
	right.initComplete()

	fmt.Println(TraverseInLayer(MergeTwoBinarys(left, right)))
}
