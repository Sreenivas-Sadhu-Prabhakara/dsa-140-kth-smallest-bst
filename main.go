package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// KthSmallest returns the k-th smallest value (1-based) in the BST; ok is
// false when k is out of range. Iterative inorder, stopped after k visits —
// the rest of the tree is never touched.
func KthSmallest(root *TreeNode, k int) (int, bool) {
	if k < 1 {
		return 0, false
	}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val, true
		}
		cur = cur.Right
	}
	return 0, false
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(5, tn(3, tn(2, tn(1, nil, nil), nil), tn(4, nil, nil)), tn(6, nil, nil))
	fmt.Println(KthSmallest(root, 3))
}
