package solution

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	printList(head)

	newHead := reverseList(head)

	printList(newHead)
}
