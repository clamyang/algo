package binarytree
// #110 平衡二叉树

import "math"

func BalancedTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := BalancedTree(root.left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := BalancedTree(root.right)
	if rightDepth == -1 {
		return -1
	}

	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return -1
	}

	return 1 + max(leftDepth, rightDepth)
}
