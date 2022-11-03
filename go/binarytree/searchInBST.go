package binarytree

// #700 二叉搜索树中搜索

// 中 左 右
func SearchInBST(root *TreeNode, target int) *TreeNode {
	if root == nil || root.val == target {
		return root
	}

	// 左
	if root.val > target {
		if data := SearchInBST(root.left, target); data != nil {
			return data
		}
	}

	// 右
	if root.val < target {
		if data := SearchInBST(root.right, target); data != nil {
			return data
		}
	}

	return nil
}

// 大佬写的递归
func SearchInBSTV2(root *TreeNode, target int) *TreeNode {
	if root == nil || root.val == target {
		return root
	}

	var data *TreeNode
	// 左
	if root.val > target {
		data = SearchInBST(root.left, target)
	}

	// 右
	if root.val < target {
		data = SearchInBST(root.right, target)
	}

	return data
}
