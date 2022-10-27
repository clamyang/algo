package binarytree

import "container/list"

// CountTreeNodesLayer 层序遍历方法最简单
func CountTreeNodesLayer(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	nodeCount := 0
	for queue.Len() != 0 {
		size := queue.Len()
		nodeCount += size
		for i := 0; i < size; i++ {
			head := queue.Remove(queue.Back()).(*TreeNode)

			if head.left != nil {
				queue.PushBack(head.left)
			}

			if head.right != nil {
				queue.PushBack(head.right)
			}
		}
	}

	return nodeCount
}

// CountTreeNodesPost
//         2
//    3         5
//  7   9
func CountTreeNodesPost(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftNodes := CountTreeNodesPost(root.left)
	rightNodes := CountTreeNodesPost(root.right)

	return 1 + leftNodes + rightNodes
}

/*
	以上方法对于普通的二叉树是完全够用了，
    如果求完全二叉树的节点个数，还有更好的方法
 	利用完全二叉树的性质
*/

// CountCompleteTreeNodes 完全二叉树求节点个数
// 求完全二叉树中满二叉树的节点个数（通过公式计算）
func CountCompleteTreeNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := root.left
	right := root.right

	leftDepth := 0
	rightDepth := 0

	for left != nil {
		leftDepth++
		left = left.left
	}

	for right != nil {
		rightDepth++
		right = right.right
	}

	if leftDepth == rightDepth {
		return (2 << leftDepth) - 1
	}

	return CountCompleteTreeNodes(root.left) + CountCompleteTreeNodes(root.right) + 1
}
