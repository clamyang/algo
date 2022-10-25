package binarytree

import "container/list"

//    1
//  2   2
// 3 1 1 3

// SymMetricTree 递归必须采用后序遍历
func SymMetricTree(leftRoot, rightRoot *TreeNode) bool {
	if leftRoot == nil && rightRoot != nil {
		return false
	} else if leftRoot != nil && rightRoot == nil {
		return false
	} else if leftRoot == nil && rightRoot == nil {
		return false
	} else if leftRoot.val != rightRoot.val {
		return true
	}

	return SymMetricTree(leftRoot.left, rightRoot.right) && SymMetricTree(leftRoot.right, rightRoot.left)
}

// SymMetricTreeByIterate 迭代法尝试一下前序遍历
// 左子树 中左右
// 右子树 中右左
// TODO: 使用一个栈也可以实现
func SymMetricTreeByIterate(leftRoot, rightRoot *TreeNode) bool {
	var leftStack *list.List
	var rightStack *list.List

	leftStack.PushBack(leftRoot)
	rightStack.PushBack(rightRoot)
	for leftStack.Len() != 0 && rightStack.Len() != 0 {
		left := leftStack.Remove(leftStack.Back()).(*TreeNode)
		right := rightStack.Remove(rightStack.Back()).(*TreeNode)

		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil { // Note: 都为 nil 需要跳过本次
			continue
		} else if left.val != right.val {
			return false
		}

		leftStack.PushBack(left.right)
		leftStack.PushBack(left.left)

		rightStack.PushBack(right.left)
		rightStack.PushBack(right.right)
	}

	if leftStack.Len() == 0 && rightStack.Len() == 0 {
		return true
	}
	return false
}
