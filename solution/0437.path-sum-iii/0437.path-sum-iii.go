package _437_path_sum_iii
//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res int

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return res
	}
	return dfs(root, sum)
}

func dfs(node *TreeNode,sum int) int {
	var tmp int
	if tmp == sum {
		res ++
	}
	if node == nil {
		tmp += 0
	}
	if node != nil {
		tmp += node.Val
	}
	if node.Left != nil {
		tmp += dfs(node.Left, sum)
	}
	if node.Right != nil {
		tmp += dfs(node.Right, sum)
	}
	return res
}