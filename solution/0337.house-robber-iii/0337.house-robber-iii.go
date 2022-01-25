package _337_house_robber_iii

//* Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//func rob(root *TreeNode) int {
//	m := make(map[*TreeNode]int)
//	return robInternal(root, m)
//}
//
//func robInternal(root *TreeNode, m map[*TreeNode]int) int {
//	if root == nil {
//		return 0
//	}
//	// 优化地方：先判断哈希缓存表中是否已有root结果，有直接返回
//	if v, ok := m[root];ok {
//		return v
//	}
//
//	money := root.Val // money：记录爷爷➕孙子偷的钱
//	if root.Left != nil {
//		money += robInternal(root.Left.Left,m) + robInternal(root.Left.Right,m)
//	}
//	if root.Right != nil {
//		money += robInternal(root.Right.Left,m) + robInternal(root.Right.Right,m)
//	}
//	// 返回结果为 Max（爷爷加孙子节点，父亲节点）
//	res := Max(money, robInternal(root.Left,m)+ robInternal(root.Right,m))
//
//	// 优化地方：将计算结果存入哈希表
//	m[root] = res
//	return res
//}

func rob(root *TreeNode) int {
	res := robInternal(root)
	// 0表示当前节点不偷，1表示当前节点偷
	return Max(res[0],res[1])
}

func robInternal(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	res := make([]int, 2)
	left := robInternal(root.Left)
	right := robInternal(root.Right)

	//当前节点选择不偷：当前节点能偷到的最大钱数 = 左孩子能偷到的钱 + 右孩子能偷到的钱
	res[0] = Max(left[0],left[1]) + Max(right[0],right[1])
	//当前节点选择偷：当前节点能偷到的最大钱数 = 当前节点的钱数 + 左孩子选择自己不偷时能得到的钱 + 右孩子选择不偷时能得到的钱
	res[1] = root.Val + left[0] + right[0]

	return res
}


func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}