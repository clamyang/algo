package binarytree

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathSum(t *testing.T) {
	root := NewTree()
	root.init()

	// 第一个节点没处理 所以传参时
	// 要将第一个节点减掉
	assert.True(t, PathSum(root, 5-root.val))
}

func TestPathSumII(t *testing.T) {
	root := NewTree()
	root.init()

	path := list.New()
	path.Init()

	path.PushBack(root.val)
	PathSumII(root, path, 5-root.val)
}
