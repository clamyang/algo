package binarytree

import (
	"container/list"
)

// ReverseBinaryTree  #226 翻转二叉树
// 迭代法实现
func ReverseBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	var ans []int
	stack.PushBack(root)

	for stack.Len() != 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, node.val)
		node.left, node.right = node.right, node.left

		if node.right != nil {
			stack.PushBack(node.right)
		}

		if node.left != nil {
			stack.PushBack(node.left)
		}

	}

	return ans
}

// ReverseByRecursive 递归实现
func ReverseByRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.left, root.right = root.right, root.left
	ReverseByRecursive(root.left)
	ReverseByRecursive(root.right)
	return root
}
