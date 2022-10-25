package binarytree

import "testing"

func TestPreTraverseMulti(t *testing.T) {
	root := NewMultiNode()
	root.init()

	PreTraverseMulti(root)
}

func TestPostTraverseMulti(t *testing.T) {
	root := NewMultiNode()
	root.init()

	PostTraverseMulti(root)
}
