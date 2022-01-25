package _543_diameter_of_binary_tree

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ans = 0
func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	if root == nil {
		return 0
	}
	maxDepth(root)
	return ans
}

// 节点node的最大深度
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lhight := maxDepth(node.Left)
	rhight := maxDepth(node.Right)
	ans = max(ans, lhight+rhight) //将每个节点最大直径(左子树深度+右子树深度)当前最大值比较并取大者
	return max(lhight, rhight)+1  //返回节点深度
}

func max(x int, y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}