package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepthPost(t *testing.T) {
	root := NewTree()
	root.init()

	assert.Equal(t, 4, MaxDepthPost(root))
}

func TestMaxDepthPre(t *testing.T) {
	root := NewTree()
	root.init()

	MaxDepthPre(root, 1)
	assert.Equal(t, 4, Ans)
}
