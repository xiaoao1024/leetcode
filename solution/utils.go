package solution

import "fmt"

func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, "->")
		p = p.Next
	}
	fmt.Println()
}
