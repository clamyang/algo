package binarytree

import  (
	"testing"
	"fmt"
)

func TestFindModeInBinaryTree(t *testing.T) {
	root := NewTree()
	root.init()

	FindModeInBinaryTree(root)
}

func TestFindModeByRecursive(t *testing.T) {
	root := NewTree()
	root.init()

	FindModeByRecursive(root)

	fmt.Println(counter)	
}
