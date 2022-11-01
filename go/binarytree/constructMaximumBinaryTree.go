package binarytree

// #654 最大二叉树

// ConstructMaximumBinaryTree [3,2,1,6,0,5]
// 题目约定，数组长度大于等于 1
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	root := new(TreeNode)
	if len(nums) == 1 {
		root.val = nums[0]

		return root
	}

	idx, maxNum := getMax(nums)
	root.val = maxNum

	leftNums, rightNums := nums[:idx], nums[idx+1:]
	if len(leftNums) > 0 {
		root.left = ConstructMaximumBinaryTree(leftNums)
	}

	if len(rightNums) > 0 {
		root.right = ConstructMaximumBinaryTree(rightNums)
	}

	return root
}

func getMax(nums []int) (int, int) {
	maxNum := 0
	index := 0
	for idx, num := range nums {
		if num > maxNum {
			maxNum = num
			index = idx
		}
	}
	return index, maxNum
}
