package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountTreeNodes(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"layer", 6},
		{"post", 6},
	}

	root := NewTree()
	root.init()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "post":
				if got := CountTreeNodesPost(root); got != tt.want {
					t.Errorf("CountTreeNodesLayer() = %v, want %v", got, tt.want)
				}

			case "layer":
				if got := CountTreeNodesLayer(root); got != tt.want {
					t.Errorf("CountTreeNodesLayer() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCountCompleteTreeNodes(t *testing.T) {
	root := NewTree()
	root.initComplete()

	assert.Equal(t, 6, CountCompleteTreeNodes(root))
}
