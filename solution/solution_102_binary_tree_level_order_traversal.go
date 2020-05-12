package solution

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		cur = []*TreeNode{root}
		res = make([][]int, 0)
	)

	for len(cur) > 0 {
		var (
			values = make([]int, 0)
			next   = make([]*TreeNode, 0)
		)

		for _, node := range cur {
			values = append(values, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		res = append(res, values)
		cur = next
	}

	return res
}
