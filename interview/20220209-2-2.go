package main

import "sort"

// 2022.02.09 杭州-心光流美-线上一面-算法题2
// 前 K 个高频元素(力扣347）
// 方法二：排序
// 时间复杂度：O(nlogn), n 是数组长度。
func topKFrequent(nums []int, k int) []int {
	var res []int
	// 1. map 统计频率
	mapNum := make(map[int]int, 0)
	// 遍历nums，记录到map中
	for _, num := range nums {
		//// map 中不存在此值，num：1
		//if value, exist := mapNum[num]; !exist {
		//	mapNum[num] = 1
		//}else {
		//	// map 中存在此值，num: value+1
		//	mapNum[num] = value + 1
		//}
		// 等价于下面一行
		mapNum[num]++
	}
	for key, _ := range mapNum {
		res = append(res, key)
	}
	// 2. 对频率排序
	sort.Slice(res, func(i, j int) bool {
		return mapNum[res[i]] > mapNum[res[j]]
	})
	// 3. 获取前K个元素
	return res[:k]
}
