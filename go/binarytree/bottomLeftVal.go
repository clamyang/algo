package binarytree

import (
	"container/list"
	"fmt"
)

// #513 二叉树左下角的值

func BottomLeftVal(root *TreeNode) {
	queue := list.New()
	queue.PushBack(root)

	data := 0
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			top := queue.Remove(queue.Front()).(*TreeNode)

			// 方法一：通过父节点进行判断
			//if top.left != nil && top.left.left == nil && top.left.right == nil {
			//	data = top.left.val
			//}

			// 只保存当前行的第一个元素
			if i == 0 {
				data = top.val
			}

			if top.left != nil {
				queue.PushBack(top.left)
			}
			if top.right != nil {
				queue.PushBack(top.right)
			}
		}
	}
	fmt.Println(data)
}
