package binarytree

import (
	"container/list"
)

// #617 合并二叉树

// 采用层序遍历的方式
func MergeTwoBinarys(leftRoot *TreeNode, rightRoot *TreeNode) *TreeNode {
	queue := list.New()

	queue.PushBack(leftRoot)
	queue.PushBack(rightRoot)

	for queue.Len() != 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)

		left.val += right.val

		if left.left != nil && right.left != nil {
			queue.PushBack(left.left)
			queue.PushBack(right.left)
		}

		if left.right != nil && right.right != nil {
			queue.PushBack(left.right)
			queue.PushBack(right.right)
		}

		if left.left == nil && right.left != nil {
			left.left = right.left
		}

		// if left.left != nil && right.left == nil {
		// 无需处理
		// }

		if left.right == nil && right.right != nil {
			left.right = right.right
		}

		// if left.right != nil && right.right == nil {
		// 无需处理
		// }

	}
	return leftRoot
}
