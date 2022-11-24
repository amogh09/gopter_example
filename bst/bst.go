package bst

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%v (%v) (%v)", t.val, t.left, t.right)
}
