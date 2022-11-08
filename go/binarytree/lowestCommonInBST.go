package binarytree

/*
		9
	5	   10
  3	  6
*/
// 自己研究的递归
func LowestCommonInBST(root, p, q *TreeNode) *TreeNode {
	if (root.val >= p.val && root.val <= q.val) || (root.val <= p.val && root.val >= q.val) {
		return root
	}

	var left *TreeNode
	if root.left != nil && root.val > p.val && root.val > q.val {
		left = LowestCommonInBST(root.left, p, q)
		if left != nil {
			return left
		}
	}

	var right *TreeNode
	if root.right != nil && root.val < p.val && root.val < q.val {
		right = LowestCommonInBST(root.right, p, q)
		if right != nil {
			return right
		}
	}

	return nil
}

// 大佬写的递归
func LowestCommonInBSTByRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.val > p.val && root.val > q.val {
		left := LowestCommonInBSTByRecursive(root.left, p, q)
		if left != nil {
			return left
		}
	}

	if root.val < p.val && root.val < q.val {
		right := LowestCommonInBSTByRecursive(root.right, p, q)
		if right != nil {
			return right
		}
	}

	return root
}

func LowestCommonInBSTByIter(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.val > p.val && root.val > q.val {
			root = root.left
		} else if root.val < p.val && root.val < q.val {
			root = root.right
		} else {
			return root
		}
	}

	return nil
}
