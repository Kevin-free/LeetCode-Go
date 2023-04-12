package main

// 2021-03-28 腾讯 技术一面 算法题

// 一个单链表，假设第一个节点我们定为下标为1，第二个为2，依次类推，
// 同时假设下标为奇数的结点是升序排序，偶数的结点是降序排序，如何让整个链表升序？
type ListNode struct {
	Val  int
	Next *ListNode
}

//7:40
func sortLinkedList(list *ListNode) *ListNode {
	res := &ListNode{}
	if list == nil {
		return res
	}
	// 1.奇偶顺序分离链表
	odd, double := divideLinkedList(list)
	// 2.偶链表升序排序
	double = reverseLinkedList(double)
	// 3.合并两个有序链表
	res = mergeTwo(odd, double)
	return res
}

//7:46
func divideLinkedList(list *ListNode) (*ListNode, *ListNode) {
	odd := list
	double := list.Next
	oddNext := &ListNode{}    // 存放奇数的下一个节点
	doubleNext := &ListNode{} // 存放偶数的下一个节点
	// deal: list{1}
	if odd.Next == nil {
		return list, nil
	}
	for odd != nil {
		oddNext = odd.Next.Next
		doubleNext = odd.Next
		odd.Next = oddNext
		if oddNext == nil {
			doubleNext.Next = nil
		} else {
			doubleNext.Next = oddNext.Next
		}
		odd = odd.Next
	}
	// 引用类型，改动 odd 即改动 list，
	return list, double
}

//7:55
func reverseLinkedList(list *ListNode) *ListNode {
	var prev *ListNode // nil node
	for list != nil {
		tmp := list.Next //
		list.Next = prev
		prev = list
		list = tmp
	}
	return list
}

//8:02
func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Val = l1.Val
			res.Next = &ListNode{}
			l1 = l1.Next
		} else {
			res.Val = l2.Val
			res.Next = &ListNode{}
			l2 = l2.Next
		}
	}
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}
	return res
}

//8:10
