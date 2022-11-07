package binarytree

// #501 二叉搜索树中的众数

import (
	"container/list"
	"fmt"
)

func FindModeInBinaryTree(root *TreeNode) {
	counter := make(map[int]int)

	stack := list.New()
	cur := root

	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			// map 的高级写法
			counter[cur.val]++
			cur = cur.right
		}
	}

	// 后续在这个 map 中取出频率最高的即可
	fmt.Println(counter)
}

var counter = make(map[int]int)

func FindModeByRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	counter[root.val]++

	FindModeByRecursive(root.left)
	FindModeByRecursive(root.right)
}
