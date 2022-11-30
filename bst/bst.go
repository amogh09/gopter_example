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

// Returns the first key in the Binary Search Tree that is greater than the provided key.
// Returns nil if there is no such key in the tree.
func (root *TreeNode) FirstKeyGreaterThan(key int) *int {
	if root == nil {
		return nil
	} else if root.key <= key {
		return root.right.FirstKeyGreaterThan(key)
	} else {
		res := root.left.FirstKeyGreaterThan(key)
		if res == nil {
			return &root.key
		} else {
			return res
		}
	}
}
