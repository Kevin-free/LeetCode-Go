package _114_flatten_binary_tree_to_linked_list

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pre *TreeNode = nil

func flatten(root *TreeNode)  {
	//pre = nil
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}

////后序遍历法
//func flatten(root *TreeNode)  {
//	if root==nil{
//		return
//	}
//	var pre *TreeNode
//	//使用2重指针
//	dfs(root,&pre)
//}
//
//func dfs(root *TreeNode,pre **TreeNode){
//	if root==nil{
//		return
//	}
//	//逆后序遍历
//	dfs(root.Right,pre)
//	dfs(root.Left,pre)
//	root.Right=(*pre)
//	root.Left=nil
//	(*pre)=root
//}