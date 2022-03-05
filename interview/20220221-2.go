package main

import "sort"

// 富途-线上一面-算法题2
// 没有重复的int数组，判断能否排列为一个等差数列？（力扣1502）
// [1,2,3,4]  true
// [1,2,4,6]  false
// [3,2,1] true

// 方法一：遍历判断法
func isValid(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	// 升序排序
	sort.Ints(nums)
	// 差值
	cha := nums[1] - nums[0]
	for i := 2; i < length; i++ {
		// 若有一个差值不同，则不为等差数列
		if nums[i]-nums[i-1] != cha {
			return false
		}
	}
	return true
}
