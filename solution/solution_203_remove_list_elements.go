package solution

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

func removeElements(head *ListNode, val int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head

	for p := tmp; p != nil && p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			q := p.Next
			for q != nil && q.Val == val {
				q = q.Next
			}

			p.Next = q
		}
	}

	return tmp.Next
}
