package main

/*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := make(map[int]bool, 0)
	p := head.Next
	tail := head

	m[head.Val] = true
	for p != nil {
		if _, ok := m[p.Val]; ok {
			p = p.Next
		} else {
			m[p.Val] = true
			tail.Next = p
			tail = p
		}
	}

	tail.Next = nil

	return head
}
