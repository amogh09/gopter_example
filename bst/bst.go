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

func (t *TreeNode) FirstKeyGreaterThan(n int) *int {
	if t == nil {
		return nil
	} else if t.val <= n {
		return t.right.FirstKeyGreaterThan(n)
	} else {
		res := t.left.FirstKeyGreaterThan(n)
		if res == nil {
			return &t.val
		} else {
			return res
		}
	}
}
