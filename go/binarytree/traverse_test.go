package binarytree

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"pre", []int{2, 5, 7, 9, 6, 3}},
		{"in", []int{9, 7, 6, 5, 2, 3}},
		{"post", []int{9, 6, 7, 5, 3, 2}},
	}

	root := NewTree()
	root.init()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "pre":
				assert.True(t, reflect.DeepEqual(tt.want, root.preByIterate()))
			case "in":
				assert.True(t, reflect.DeepEqual(tt.want, root.inByIterate()))
			case "post":
				assert.True(t, reflect.DeepEqual(tt.want, root.postByIterate()))
			}
		})
	}
}
