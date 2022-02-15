package main

// 2022.02.09 杭州-心光流美-线上一面-算法题1
// 二分查找
// nums中查找 target，有则返回其下标，无则返回-1
func binarySearch(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		// 防止溢出
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	// 遍历完没找到，返回-1
	return -1
}
