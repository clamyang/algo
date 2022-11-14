package binarytree

import (
	"fmt"
	"testing"
)

func TestConvertSortedArrToBST(t *testing.T) {
	sortedArr := []int{-1, 3, 6, 7, 9, 12}
	fmt.Println(TraverseInLayer(ConvertSortedArrToBST(sortedArr)))
}
