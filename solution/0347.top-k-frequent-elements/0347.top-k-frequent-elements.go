package topkfrequentelements

// [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
// tag：
// middle
func topKFrequent(nums []int, k int) []int {
	var res []int
	// 1.使用Map记录元素值和频率，map[num]count
	mapNumCount := make(map[int]int, 0)
	for _, num := range nums {
		mapNumCount[num]++
	}

	// 2.使用小根堆获取频率前 k 高的元素
}

type Item struct {
	num   int //元素
	count int //次数
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	pq = append(pq, item)
}

func (pq PriorityQueue) Pop() interface{} {
	n := len(pq)
	item := pq[n-1]
	pq = pq[:n-1]
	return item
}
