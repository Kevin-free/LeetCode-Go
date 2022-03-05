package main

//
//import "sort"
//
//// 一个单链表，假设第一个节点我们定为下标为1，第二个为2，依次类推，
//// 同时假设下标为奇数的结点是升序排序，偶数的结点是降序排序，如何让整个链表升序？
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
////21:48
//func sortLinkedList(list *ListNode) *ListNode {
//	// 1.分离奇偶序列
//	i := 1
//	var oddList, doubleList []int
//	for list.Next != nil {
//		if i % 2 == 1 {
//			oddList = append(oddList, list.Val)
//		}else {
//			doubleList = append(doubleList, list.Val)
//		}
//		i++
//		list = list.Next
//	}
//
//	// 2.偶序列升序排序
//	sort.Ints(doubleList)
//
//	// 3.合并两个有序序列为一个
//	var resList []int
//	for i := 0; i < len(oddList); i++ {
//		for j := 0; j < len(doubleList); j++ {
//			if oddList[i] < doubleList[j] {
//				resList = append(resList, oddList[i])
//				i++
//			}else {
//				resList = append(resList, doubleList[j])
//				j++
//			}
//		}
//	}
//
//	var res *ListNode
//	for _, val := range resList {
//		res.Val = val
//
//	}
//
//	return &ListNode{}
//}
