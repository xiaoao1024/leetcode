package solution

func reverseList(head *ListNode) *ListNode {
	var (
		newHead = &ListNode{}
		p       *ListNode
	)

	for head != nil {
		p = head
		head = head.Next
		p.Next = newHead.Next
		newHead.Next = p
	}

	return newHead.Next
}
