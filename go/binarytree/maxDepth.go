package binarytree

import "container/list"

/*
 * 二叉树深度：指从根节点到该节点的最长简单路径边的条数或者节点数（取决于深度从0开始还是从1开始）
 *
 * 二叉树高度：指从该节点到叶子节点的最长简单路径边的条数或者节点数（取决于高度从0开始还是从1开始）
 *
 **/

// MaxDepthPost 后序遍历
func MaxDepthPost(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepthPost(root.left)
	rightDepth := MaxDepthPost(root.right)

	return 1 + max(leftDepth, rightDepth)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var Ans int

// MaxDepthPre 前序遍历
func MaxDepthPre(root *TreeNode, depth int) {
	if depth > Ans {
		Ans = depth
	}

	if root.left == nil && root.right == nil {
		return
	}

	if root.left != nil {
		depth++
		MaxDepthPre(root.left, depth)
		depth--
	}

	if root.right != nil {
		depth++
		MaxDepthPre(root.right, depth)
		depth--
	}
}

// MaxDepthIterate 迭代法
func MaxDepthIterate(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	var layer int
	for queue.Len() != 0 {
		size := queue.Len()
		layer++
		for i := 0; i < size; i++ {
			head := queue.Remove(queue.Front()).(*TreeNode)

			if head.left != nil {
				queue.PushBack(head.left)
			}

			if head.right != nil {
				queue.PushBack(head.right)
			}
		}
	}
}
