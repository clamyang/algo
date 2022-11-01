package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 0, 5, 6}, ConstructMaximumBinaryTree(([]int{3, 2, 1, 6, 0, 5})).postByIterate())
}
