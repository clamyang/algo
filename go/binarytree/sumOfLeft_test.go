package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfLeft(t *testing.T) {
	root := NewTree()
	root.init()
	res := new(int)
	SumOfLeft(root, res)
	assert.Equal(t, 9, *res)
}

func TestSumOfLeftV2(t *testing.T) {
	root := NewTree()
	root.init()
	assert.Equal(t, 9, SumOfLeftV2(root))
}
