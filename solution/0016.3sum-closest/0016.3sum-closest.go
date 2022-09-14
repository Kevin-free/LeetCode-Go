package _016_3sum_closest

import (
	"LeetCode-Go/util"
	"sort"
)

/**
16. 最接近的三数之和
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // 排序
	length := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < length; i++ {
		l, r := i+1, length-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if util.AbsInt(target-sum) < util.AbsInt(target-ans) {
				ans = sum
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return ans
			}
		}
	}
	return ans
}
