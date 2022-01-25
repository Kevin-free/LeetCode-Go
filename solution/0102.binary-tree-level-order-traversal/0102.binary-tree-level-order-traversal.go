package _102_binary_tree_level_order_traversal

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root) // 添加到队尾，等于add()
	for len(queue) != 0 {
		count := len(queue)
		var list []int
		for count > 0 {
			node := queue[0] //取出队首，等于poll()
			queue = queue[1:] //移出队首更新队列
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count--
		}
		res = append(res, list)
	}
	return res
}
