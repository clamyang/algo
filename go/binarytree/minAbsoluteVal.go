package binarytree

// #530 minimum-absolute-difference-in-bst

import "container/list"

func MinAbsoluteValByLayer(root *TreeNode) int {
	return getByLayer(root)
}

// NOTE: 这里严重错误，写出的答案与题意不符..
// 人家求的是绝对值的最小值，结果自己用层序求了最大值..
// 确实傻呗...
func getByLayer(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	minVal, maxVal := 0, 0
	getMin, getMax := false, false

	if root.left != nil && root.right == nil {
		maxVal = root.val
		getMax = true
	}
	if root.left == nil && root.right != nil {
		minVal = root.val
		getMin = true
	}

	if root.left == nil && root.right == nil {
		return root.val
	}
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			head := queue.Remove(queue.Front()).(*TreeNode)
			if !getMax && head.left == nil && head.right == nil {
				getMax = true
				maxVal = head.val
			}
			if !getMin && size-1 == i {
				minVal = head.val
			}
			if head.right != nil {
				queue.PushBack(head.right)
			}
			if head.left != nil {
				queue.PushBack(head.left)
			}
		}
	}

	return maxVal - minVal
}

func MinAbsoluteVal(root *TreeNode) {
	minVal := 100

	stack := list.New()

	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil {
				minVal = min(minVal, cur.val-pre.val)
			}

			pre = cur

			cur = cur.right
		}
	}
}
