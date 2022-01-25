package _538_convert_bst_to_greater_tree

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num = 0

func convertBST(root *TreeNode) *TreeNode {
	//num =0
	if root != nil {
		convertBST(root.Right)
		root.Val = root.Val + num
		num = root.Val
		convertBST(root.Left)
		return root
	}
	return nil
}
