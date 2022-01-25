package _104_maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lhight := maxDepth(root.Left)
	rhight := maxDepth(root.Right)
	return max(lhight, rhight) + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
