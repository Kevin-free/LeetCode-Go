package _108_convert_sorted_array_to_binary_search_tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1 // go 中的位运算符优先级高于加减，必须带括号！
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
