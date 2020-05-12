package solution

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(levelOrder(root))
}
