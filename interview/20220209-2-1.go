package main

import "container/heap"

// 2022.02.09 杭州-心光流美-线上一面-算法题2
// 前 K 个高频元素(力扣347）
// 方法一：小顶堆
func topKFrequent(nums []int, k int) []int {
	var res []int
	// 1. map 统计频率
	mapNum := make(map[int]int, 0)
	// 遍历nums，记录到map中
	for _, num := range nums {
		mapNum[num]++
	}
	// 2. 优先队列 对频率排序
	pq := PriorityQueue{}
	for num, count := range mapNum {
		heap.Push(&pq, &Item{num: num, count: count})
	}
	// 3. 获取前K个元素
	for len(res) < k {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.num)
	}
	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	num   int // 元素
	count int // 次数
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们想要 Pop 出最大值，因为 golang 中 heap 是最小堆，所以用大于判断
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
