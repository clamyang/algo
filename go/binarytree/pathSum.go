package binarytree

import (
	"container/list"
)

// #112 path sum
/*
  递归函数什么时候需要返回值：
	1.如果需要搜索整棵二叉树且不用处理递归返回值，递归函数就不要返回值。
	2.如果需要搜索整棵二叉树且需要处理递归返回值，递归函数就需要返回值。
      （这里判断是否需要处理的依据是？是否根据左子树的返回值和右子树的返回值进行处理，参考最大最小深度
		最大深度时，需要根据左右子树的返回值进行判断，然后返回给父节点）
	3.如果要搜索其中一条符合条件的路径，那么递归一定需要返回值，因为遇到符合条件的路径了就要及时返回。
*/

func PathSum(root *TreeNode, count int) bool {
	// 采用减法的方式代码更简洁
	if root.left == nil && root.right == nil && count == 0 {
		return true
	}
	if root.left == nil && root.right == nil {
		return false
	}

	// 前序遍历，中后都可以，因为都没有对 中 这个节点进行处理
	// 中 不做处理

	// 左
	if root.left != nil {
		// 省略了回溯过程，把 count-root.left.val 直接当作参数传入
		if PathSum(root.left, count-root.left.val) {
			return true
		}
	}

	// 右
	if root.right != nil {
		if PathSum(root.right, count-root.right.val) {
			return true
		}
	}

	return false
}

func PathSumII(root *TreeNode, path *list.List, count int) {
	if root.left == nil && root.right == nil && count == 0 {
		cur := path.Front()
		for cur != nil {
			println(cur.Value.(int))
			cur = cur.Next()
		}

		return
	}

	if count < 0 {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}

	if root.left != nil {
		path.PushBack(root.left.val)
		PathSumII(root.left, path, count-root.left.val)
		path.Remove(path.Back())
	}

	if root.right != nil {
		path.PushBack(root.right.val)
		PathSumII(root.right, path, count-root.right.val)
		path.Remove(path.Back())
	}
}
