package binarytree
// #257 所有路径

import (
	"fmt"
	"strconv"
)

// #257 二叉树所有路径

// TreePaths 前序遍历 中 左 右
func TreePaths(root *TreeNode, paths []int) {
	// 中
	paths = append(paths, root.val)
	if root.left == nil && root.right == nil {
		single := ""
		for _, node := range paths {
			single += strconv.Itoa(node)
			single += "->"
		}
		fmt.Println(single)
	}

	// 左
	if root.left != nil {
		TreePaths(root.left, paths)
		// 回溯
		remove(paths)
	}

	// 右
	if root.right != nil {
		TreePaths(root.right, paths)
		// 回溯
		remove(paths)
	}
}

func remove(paths []int) {
	if len(paths) == 0 {
		return
	}

	paths = paths[:len(paths)-1]
}
