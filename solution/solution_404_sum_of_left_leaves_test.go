package solution

import (
	"testing"
)

var root = &TreeNode{
	Val:  3,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	},
}

func Test_sumOfLeftLeaves(t *testing.T) {
	sum := sumOfLeftLeaves(root)
	if sum != 24 {
		t.Errorf("expect 24, actual %d", sum)
	}
}
