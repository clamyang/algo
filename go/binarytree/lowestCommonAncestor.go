package binarytree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := LowestCommonAncestor(root.left, p, q)
	right := LowestCommonAncestor(root.right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	}

	return nil
}
