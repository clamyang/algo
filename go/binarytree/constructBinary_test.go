package binarytree

import (
	"fmt"
	"testing"
)

func TestConstructBinaryTreeFromInAndPost(t *testing.T) {
	root := ConstructBinaryTreeFromInAndPost([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(TraverseInLayer(root))
}

func TestConstructBinaryTreeFromPreAndIn(t *testing.T) {
	root := ConstructBinaryTreeFromPreAndIn([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(TraverseInLayer(root))
}
