package binarytree

// #669 修剪二叉搜索树

func TrimBST(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.val < low {
		return TrimBST(root.right, low, high)
	}

	if root.val > high {
		return TrimBST(root.left, low, high)
	}

	root.left = TrimBST(root.left, low, high)
	root.right = TrimBST(root.right, low, high)

	return root
}
