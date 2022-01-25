package _215_kth_largest_element_in_an_array

import (
	"container/heap"
)

type TopList []int

func (t TopList) Len() int {
	return len(t)
}

func (t TopList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t TopList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TopList) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *TopList) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	m := make(TopList, 0)
	size := 0
	for i := range nums {
		if size < k {
			heap.Push(&m, nums[i])
			size++
		} else {
			if m[0] < nums[i] { //小顶堆 堆顶元素小于当前元素
				heap.Pop(&m)
				heap.Push(&m, nums[i])
			}
		}
	}
	return m[0]
}


//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums) - k]
//}