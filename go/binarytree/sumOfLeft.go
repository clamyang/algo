package binarytree

func SumOfLeft(root *TreeNode, sum *int) {
	// 找到左节点
	if root.left != nil && root.left.left == nil && root.left.right == nil {
		*sum += root.left.val
	}

	if root.left != nil {
		SumOfLeft(root.left, sum)
	}

	if root.right != nil {
		SumOfLeft(root.right, sum)
	}
}

func SumOfLeftV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := SumOfLeftV2(root.left)
	if root.left != nil && root.left.left == nil && root.left.right == nil {
		left = root.left.val
	}

	right := SumOfLeftV2(root.right)

	return left + right
}
