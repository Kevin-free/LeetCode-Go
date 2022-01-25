package _209_minimum_size_subarray_sum

import (
	"math"
	"myutil"
)

/**
https://leetcode-cn.com/problems/minimum-size-subarray-sum/
209. 长度最小的子数组
*/

// 暴力
//func minSubArrayLen(s int, nums []int) int {
//	length := len(nums)
//	ans := length + 1
//	for start := 0; start < length; start++ {
//		sum := nums[start]
//		if sum >= s {
//			return 1
//		} else {
//			for end := start + 1; end < length; {
//				sum += nums[end]
//				if sum >= s {
//					len := end - start + 1
//					ans = myutil.MinInt(ans, len)
//					break
//				} else {
//					end++
//				}
//			}
//
//		}
//	}
//	if ans > length {
//		return 0
//	}
//	return ans
//}

// 双指针
func minSubArrayLen(s int, nums []int) int {
	length := len(nums) // 内置函数得出数组长度
	if length == 0 {
		return 0
	}
	ans := math.MaxInt32 // 定义结果为最大值
	start, end := 0, 0 // 左右边界
	sum := 0 // 临时变量存和
	// 遍历数组
	for end < length {
		// 求和
		sum += nums[end]
		// 当 sum 大或等于目标 s 时，则更新子数组的最小长度（此时子数组的长度是 end-start+1），并减去 nums[start] ，将 start 右移
		for sum >= s {
			ans = myutil.MinInt(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		// sum 小于 s，将 end 右移
		end++
	}
	// 若果 ans 未改动，无符号条件返回 0
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
