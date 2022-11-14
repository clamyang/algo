package binarytree

// #108 从有序数组构建高度平衡二叉树

func ConvertSortedArrToBST(sortedArr []int) *TreeNode {
	if len(sortedArr) == 0 {
		return nil
	}

	mid, index := getMiddle(sortedArr)
	newNode := new(TreeNode)
	newNode.val = mid

	newNode.left = ConvertSortedArrToBST(sortedArr[:index])
	newNode.right = ConvertSortedArrToBST(sortedArr[index+1:])

	return newNode
}

func getMiddle(sortedArr []int) (int, int) {
	length := len(sortedArr)

	return sortedArr[length/2], length / 2
}
