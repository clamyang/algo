package binarytree

// #450 删除二叉搜索树中的节点

func DeleteNodeInBST(root *TreeNode, target int) *TreeNode {
	// 第一种情况：没有找到
	if root == nil {
		return root
	}

	if root.val == target {
		// 第二种情况，左右子树都为空，叶子节点
		if root.left == nil && root.right == nil {
			root = nil
			return root
		}

		// 第三种情况，左子树不为空，右子树为空
		if root.left == nil {
			root = root.right 
			return root
		}

		// 第四种情况，左子树为空，右子树不为空
		if root.right == nil {
			root = root.left
			return root
		}

		// 第五种情况，左子树不为空，右子树不为空
		if root.left != nil && root.right != nil {
			subRoot := root.right
			for subRoot.left != nil {
				subRoot = subRoot.left
			}
			
			subRoot.left = root.left
			root = root.right

			return root
		}
	}

	if root.val > target {
		root.left = DeleteNodeInBST(root.left, target)
	}

	if root.val < target {
		root.right = DeleteNodeInBST(root.right, target)
	}

	return root
}
