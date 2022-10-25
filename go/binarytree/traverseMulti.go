package binarytree

import "fmt"

type MultiNode struct {
	val    int
	childs []*MultiNode
}

func NewMultiNode() *MultiNode {
	return new(MultiNode)
}

/*
	         2
    3        5       8
12 23 56     21    34 67
*/
func (node *MultiNode) init() {
	node.val = 2
	node.childs = make([]*MultiNode, 3)
	node.childs[0] = &MultiNode{3, nil}
	node.childs[1] = &MultiNode{5, nil}
	node.childs[2] = &MultiNode{8, nil}

	node.childs[0].childs = make([]*MultiNode, 3)
	node.childs[1].childs = make([]*MultiNode, 1)
	node.childs[2].childs = make([]*MultiNode, 2)

	node.childs[0].childs[0] = &MultiNode{12, nil}
	node.childs[0].childs[1] = &MultiNode{23, nil}
	node.childs[0].childs[2] = &MultiNode{65, nil}

	node.childs[1].childs[0] = &MultiNode{21, nil}
	node.childs[2].childs[0] = &MultiNode{34, nil}
	node.childs[2].childs[1] = &MultiNode{67, nil}
}

// PreTraverseMulti N叉树的前序遍历
func PreTraverseMulti(root *MultiNode) {
	if root == nil {
		return
	}

	fmt.Println(root.val)
	for _, child := range root.childs {
		PreTraverseMulti(child)
	}
}

// PostTraverseMulti #590 N叉树的后续遍历
func PostTraverseMulti(root *MultiNode) {
	if root == nil {
		return
	}

	for _, child := range root.childs {
		PostTraverseMulti(child)
	}
	fmt.Println(root.val)
}
