package binarytree

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTree() *TreeNode {
	return new(TreeNode)
}

func (node *TreeNode) init() {
	node.val = 2

	node.left = new(TreeNode)
	node.left.val = 5

	node.right = new(TreeNode)
	node.right.val = 3

	node.left.left = new(TreeNode)
	node.left.left.val = 7

	node.left.left.left = new(TreeNode)
	node.left.left.left.val = 9

	node.left.left.right = new(TreeNode)
	node.left.left.right.val = 6
}

func (node *TreeNode) initComplete() {
	node.val = 2

	node.left = new(TreeNode)
	node.left.val = 5

	node.right = new(TreeNode)
	node.right.val = 3

	node.left.left = new(TreeNode)
	node.left.left.val = 7
	node.left.right = new(TreeNode)
	node.left.right.val = 8

	node.left.left.left = new(TreeNode)
	node.left.left.left.val = 9
}

//    6
//  5   7
// 3 4
//2
func (node *TreeNode) initSearchTree() {
	node.val = 7

	node.left = new(TreeNode)
	node.left.val = 4

	node.right = new(TreeNode)
	node.right.val = 8

	node.left.left = new(TreeNode)
	node.left.left.val = 3
	node.left.right = new(TreeNode)
	node.left.right.val = 6

	node.left.left.left = new(TreeNode)
	node.left.left.left.val = 2
}

// 中 左 右
func (node *TreeNode) preOrder() {
	if node == nil {
		return
	}

	fmt.Printf("%2d", node.val)
	node.left.preOrder()
	node.right.preOrder()
}

func (node *TreeNode) intOrder() {
	if node == nil {
		return
	}

	node.left.intOrder()
	fmt.Printf("%2d", node.val)
	node.right.intOrder()
}

func (node *TreeNode) postOrder() {
	if node == nil {
		return
	}

	node.left.postOrder()
	node.right.postOrder()
	fmt.Printf("%2d", node.val)
}

func (node *TreeNode) preByIterate() []int {
	stack := list.New()

	var ans []int

	stack.PushBack(node)
	for stack.Len() != 0 {
		top := stack.Remove(stack.Back()).(*TreeNode)

		ans = append(ans, top.val)

		if top.right != nil {
			stack.PushBack(top.right)
		}

		if top.left != nil {
			stack.PushBack(top.left)
		}
	}
	return ans
}

func (node *TreeNode) inByIterate() []int {
	cur := node
	stack := list.New()

	var ans []int

	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			ans = append(ans, cur.val)
			cur = cur.right
		}
	}

	return ans
}

func (node *TreeNode) postByIterate() []int {
	var ans []int
	if node == nil {
		return ans
	}

	st := list.New()
	st.PushBack(node)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.val)
		if node.left != nil {
			st.PushBack(node.left)
		}
		if node.right != nil {
			st.PushBack(node.right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
