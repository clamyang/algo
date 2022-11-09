package binarytree

// #701 二叉搜索树中的插入操作

/*
		7
	4		8
3 		6
2
*/
func InsertToBSTByIter(root *TreeNode, target int) *TreeNode {
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
