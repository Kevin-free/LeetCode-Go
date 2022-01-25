package _019_remove_nth_node_from_end_of_list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
// 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	// 找到要删除的节点（倒数第 n+1 个）
	del := findFromEnd(dummy, n+1)
	// 当删除第一个节点时，此写法空指针异常
	//del := findFromEnd(head, n+1)
	del.Next = del.Next.Next
	return dummy.Next
}

// 返回链表的倒数第 n 个节点
func findFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
