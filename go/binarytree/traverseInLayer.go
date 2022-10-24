package binarytree

import (
	"container/list"
)

func TraverseInLayer(root *TreeNode) [][]int {
	queue := list.New()

	var ans [][]int

	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()

		var layer []int
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			layer = append(layer, node.val)

			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
		ans = append(ans, layer)
	}

	return ans
}
