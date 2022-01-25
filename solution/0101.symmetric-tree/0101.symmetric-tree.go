package _101_symmetric_tree


 //* Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}else{
		return partSym(root.Left, root.Right)
	}
}

func partSym(left, right *TreeNode) bool {
	//if left.Left == nil ||  left.Right == nil || right.Left == nil ||  right.Right == nil{
	//	return false
	//}
	if (left.Val == right.Val) && partSym(left.Left, right.Right) && partSym(left.Right, right.Left) {
		return true
	}else {
		return false
	}
}