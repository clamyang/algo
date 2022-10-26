package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDepth(t *testing.T) {
	root := NewTree()
	root.init()

	assert.Equal(t, 2, MinDepth(root))
}

func TestMinDepthIterate(t *testing.T) {
	root := NewTree()
	root.init()

	assert.Equal(t, 2, MinDepthIterate(root))

}

func TestMinDepthPre(t *testing.T) {
	root := NewTree()
	root.init()

	MinDepthPre(root, 1)
	assert.Equal(t, 2, minDepth)
}
