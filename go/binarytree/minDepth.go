package binarytree

// 最小深度：是根节点到最近叶子节点的最短路径上的节点数量。

// MinDepth 注意：是从根节点到 **叶子结点**
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
