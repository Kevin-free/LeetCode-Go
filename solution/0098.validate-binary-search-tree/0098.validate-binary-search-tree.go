package _098_validate_binary_search_tree

import "math"

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pre = math.MinInt64
func isValidBST(root *TreeNode) bool {
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := dfs(root.Left)
	if pre < root.Val {
		pre = root.Val
	}else {
		return false
	}
	right := dfs(root.Right)
	return left && right
}

