package _094_binary_tree_inorder_traversal

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var res []int

//func inorderTraversal(root *TreeNode) []int {
//    res = make([]int, 0)
//    inorder(root)
//    return res
//}
//
//func inorder(root *TreeNode) {
//    if root != nil {
//        inorder(root.Left)
//        res = append(res, root.Val)
//        inorder(root.Right)
//    }
//}

//func inorderTraversal(root *TreeNode) []int {
//	var res []int
//	var stack []*TreeNode
//
//	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
//		for root != nil {
//			stack = append(stack, root) //入栈
//			root = root.Left            //移至最左
//		}
//		index := len(stack) - 1             //栈顶
//		res = append(res, stack[index].Val) //中序输出
//		root = stack[index].Right           //右节点会进入下次循环，如果 =nil，继续出栈
//		stack = stack[:index]               //出栈
//	}
//	return res
//}

type ColorNode struct {
	node *TreeNode
	color string
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*ColorNode
	stack = append(stack, &ColorNode{root, "white"})
	var cn *ColorNode

	for len(stack) != 0 {
		cn = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 以上两句等同于 cn = stack.pop() ，别忘了加这句
		if cn.color == "white" {
			// 因为栈是先进后出，所以中序是 右-根-左 的顺序添加
			if cn.node.Right != nil {
				stack = append(stack, &ColorNode{cn.node.Right,"white"})
			}
			stack = append(stack,&ColorNode{cn.node, "gray"})
			if cn.node.Left != nil {
				stack = append(stack, &ColorNode{cn.node.Left, "white"})
			}
		}else {
			res = append(res, cn.node.Val)
		}
	}

	return res
}