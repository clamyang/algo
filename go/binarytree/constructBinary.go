package binarytree

// #106 从中序和后序构建二叉树

// ConstructBinaryTreeFromInAndPost
// inorder = [9,3,15,20,7] postorder = [9,15,7,20,3]
func ConstructBinaryTreeFromInAndPost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	middle := postorder[len(postorder)-1]

	root := new(TreeNode)
	root.val = middle

	if len(postorder) == 1 {
		return root
	}

	inMiddle := getIndex(middle, inorder)
	leftIn, rightIn := inorder[:inMiddle], inorder[inMiddle+1:]

	postorder = postorder[:len(postorder)-1]
	leftPost, rightPost := postorder[:len(leftIn)], postorder[len(leftIn):]

	root.left = ConstructBinaryTreeFromInAndPost(leftIn, leftPost)
	root.right = ConstructBinaryTreeFromInAndPost(rightIn, rightPost)

	return root
}

func getIndex(val int, arr []int) int {
	for idx, data := range arr {
		if data == val {
			return idx
		}
	}
	return 0
}

// ConstructBinaryTreeFromPreAndIn 从前序和中序构造二叉树
func ConstructBinaryTreeFromPreAndIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	middle := preorder[0]
	root.val = middle

	if len(preorder) == 1 {
		return root
	}

	inIdx := getIndex(middle, inorder)

	leftIn, rightIn := inorder[:inIdx], inorder[inIdx+1:]

	preorder = preorder[1:]

	leftpre, rightpre := preorder[:len(leftIn)], preorder[len(leftIn):]

	root.left = ConstructBinaryTreeFromPreAndIn(leftpre, leftIn)
	root.right = ConstructBinaryTreeFromPreAndIn(rightpre, rightIn)

	return root
}
