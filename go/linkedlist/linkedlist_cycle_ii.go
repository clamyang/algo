package linkedlist

// #142 寻找环的入口

type Node struct {
	Val  int
	Next *Node
}

// 2*(x+y) = x+y+n(y+z)
// x+y = n*(y+z)
// x = (n-1)(y+z)+z ---> x = z
func detectCycle(head *Node) *Node {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}

			return head
		}
	}

	return nil
}
