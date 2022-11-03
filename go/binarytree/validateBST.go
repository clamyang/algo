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
// 需要理解 pre 指针的含义，二叉搜索树的中序遍历结果是递增的，
// 使用 pre 指向前一个节点，如果前一个节点值大于当前节点，就不是二叉搜索树。
func IsValidBST(root *TreeNode) bool {
	stack := list.New()
	stack.PushBack(root)

	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil && cur.val <= pre.val {
				return false
			}

			pre = cur

			cur = cur.right
		}
	}

	return true
}
