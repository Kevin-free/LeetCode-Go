package _206_reverse_linked_list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. 反转链表:简单
// tags：链表，递归

// 迭代写法：时间复杂度 O（n)，空间复杂度 O(1)
func reverseList(head *ListNode) *ListNode {
 // 初始化 prev 为空节点
	var prev *ListNode
// 遍历整条链表
	for head != nil {
		//反转链表，注意顺序！
		next := head.Next // ① 先存下一个节点，不然后面找不着
		head.Next = prev  // ② 当前节点指向上一个节点
		prev = head // ③ 向后移动一个节点
		head = next // ④ 向后移动一个节点
		// 语法性质可简写如下
		// head.Next, head, prev = prev, head.Next, head
	}
	// 此时 prev 指向原链表的最后一个节点
	return prev
}

// 递归写法：时间复杂度 O（n)，空间复杂度 O(n) 递归解法需要堆栈
func reverseList(head *ListNode) *ListNode {
	// 递归终止条件
   if head == nil || head.Next == nil {
       return head
   }
   // 递归调用
   last := reverseList(head.Next)
   // 每层递归中的处理
   head.Next.Next = head // 核心：反转链表
   head.Next = nil // 防止链表循环，需要将head.next设置为空
   // last 指向原链表的最后一个节点
   return last
}