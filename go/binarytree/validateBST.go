package binarytree

import (
	"container/list"

)

// #98 验证二叉搜索树

// 关键需要知道中序遍历的二叉搜索树是有序数组
// 所以可以通过中序遍历将树转化为数组验证数组是否有序即可

type inOrderArr []int

func (in inOrderArr) Len() int {
	return len(in)
}

func (in inOrderArr) Less(i, j int) bool {
	return in[i] < in[j]
}

func (in inOrderArr) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

var arrIn = make(inOrderArr, 0, 10)

func ValidateBSTByIn(root *TreeNode) {

	if root.left != nil {
		ValidateBSTByIn(root.left)
	}
	arrIn = append(arrIn, root.val)

	if root.right != nil {
		ValidateBSTByIn(root.right)
	}
}

// 中序迭代法方式实现
func IsValidBST(root *TreeNode) bool {
	stack := list.New()
	stack.PushBack(root)

	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() != 0 {
		if cur.left != nil {
			stack.PushBack(cur.left)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil && pre.val >= cur.val {
				return false
			}

			pre = cur

			cur = cur.right
		}
	}

	return true
}
