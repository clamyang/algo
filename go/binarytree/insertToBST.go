package binarytree

// #701 二叉搜索树中的插入操作

/*
		7
	4		8
3 		6
2
*/
func InsertIntoBSTByIter(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return &TreeNode{val: target}
	}

	cur, pre := root, new(TreeNode)

	for cur != nil {
		pre = cur
		if target > cur.val {
			cur = cur.right
			if cur == nil {
				pre.right = &TreeNode{val: target}
			}
		} else if target < cur.val {
			cur = cur.left
			if cur == nil {
				pre.left = &TreeNode{val: target}
			}
		}
	}

	return root
}

// 在递归的过程中给子节点赋值
func InsertIntoBSTByRecursive(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return &TreeNode{val: target}
	}

	if root.val > target {
		root.left  = InsertIntoBSTByRecursive(root.left, target)
	}

	if root.val < target {
		root.right = InsertIntoBSTByRecursive(root.right, target)
	}

	return root
}
