package binarytree

import (
	"container/list"
	"math"
)

// 最小深度：是根节点到最近叶子节点的最短路径上的节点数量。

// MinDepth 注意：是从根节点到 **叶子结点**
//         2
//    3
//  7   9
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MinDepth(root.left)
	rightDepth := MinDepth(root.right)

	if root.left == nil && root.right != nil {
		return 1 + rightDepth
	}

	if root.right == nil && root.left != nil {
		return 1 + leftDepth
	}

	return 1 + min(leftDepth, rightDepth)
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// minDepth
var minDepth = math.MaxInt

// MinDepthPre 前序遍历
//         2
//    3         5
//  7   9
func MinDepthPre(root *TreeNode, depth int) {
	if root.left == nil && root.right == nil {
		minDepth = min(minDepth, depth)
		return
	}

	// 中不进行处理

	// 左
	if root.left != nil {
		MinDepthPre(root.left, depth+1)
	}

	// 右
	if root.right != nil {
		MinDepthPre(root.right, depth+1)
	}

	return
}

// MinDepthIterate 迭代法是按照每一层从左到右的顺序进行遍历的
// 如果遇到一个左右孩子都为空的节点说明就到了最低层
func MinDepthIterate(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	layer := 0
	for queue.Len() != 0 {
		size := queue.Len()
		layer++

		for i := 0; i < size; i++ {
			head := queue.Remove(queue.Back()).(*TreeNode)

			if head.left != nil {
				queue.PushBack(head.left)
			}
			if head.right != nil {
				queue.PushBack(head.right)
			}

			if head.left == nil && head.right == nil {
				return layer
			}
		}
	}

	return layer
}
