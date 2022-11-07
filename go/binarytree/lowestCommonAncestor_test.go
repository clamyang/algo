package binarytree

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	root := NewTree()
	root.initSearchTree()

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1", args{root, root.left, root.right}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q)
			fmt.Println(got)

		})
	}
}
