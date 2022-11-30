package bst

import "fmt"

type TreeNode struct {
	key   int
	left  *TreeNode
	right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%v (%v) (%v)", t.key, t.left, t.right)
}

func (t *TreeNode) FirstKeyGreaterThan(key int) *int {
	if t == nil {
		return nil
	} else if t.key <= key {
		return t.right.FirstKeyGreaterThan(key)
	} else {
		res := t.left.FirstKeyGreaterThan(key)
		if res == nil {
			return &t.key
		} else {
			return res
		}
	}
}
