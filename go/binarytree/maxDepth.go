package binarytree

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

// MaxDepthPre 前序遍历
func MaxDepthPre(root *TreeNode) int {
	return 0
}
