package binarytree

// #538 二叉搜索树转换为累加树

// 最简单的方式，利用二叉搜索树中序遍历为递增的特性
// 构造出累加后的数组，根据这个数组进行构造，但是构造出来的数组是递减的
// 所以需要从后往前进行处理

var preVal int

func ConvertBSTToGreaterTree(root *TreeNode) {
	if root == nil {
		return
	}

	ConvertBSTToGreaterTree(root.right)
	root.val += preVal
	preVal = root.val
	ConvertBSTToGreaterTree(root.left)
}
