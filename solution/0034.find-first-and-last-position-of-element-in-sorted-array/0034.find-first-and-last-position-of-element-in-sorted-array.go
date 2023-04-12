package _034_find_first_and_last_position_of_element_in_sorted_array

// import "sort"

// // 34. 在排序数组中查找元素的第一个和最后一个位置：中等
// // tag：数组、二分查找
// // 只用一个边界方法
// func searchRangeOneBound(nums []int, target int) []int {
// 	// 目标值开始位置：为目标值的左侧边界，无此值则返回比它大的数的左侧边界
// 	start := findLeftBound(nums, target)
// 	// 如果开始位置越界 或 目标值不存在，直接返回
// 	if start == len(nums) || nums[start] != target {
// 		return []int{-1,-1}
// 	}
// 	// 目标值结束位置：为目标值+1的左侧边界-1，无此值则返回比它大的数的左侧边界-1
// 	end := findLeftBound(nums, target+1) - 1
// 	return []int{start, end}
// }

// // 用两个边界方法
// func searchRangeTwoBound(nums []int, target int) []int {
// 	// 目标值开始位置：为 target 的左侧边界
// 	start := findLeftBound(nums, target)
// 	// 如果开始位置越界 或 目标值不存在，直接返回
// 	if start == len(nums) || nums[start] != target {
// 		return []int{-1,-1}
// 	}
// 	// 目标值结束位置：为 target 的右侧边界
// 	end := findRightBound(nums, target)
// 	return []int{start, end}
// }

// // 用go实现的二分查找方法
// func searchRange(nums []int, target int) []int {
// 	// sort.SearchInts 为 go 实现的二分查找
// 	start := sort.SearchInts(nums, target)
// 	if start == len(nums) || nums[start] != target {
// 		return []int{-1, -1}
// 	}
// 	end := sort.SearchInts(nums, target + 1) - 1
// 	return []int{start, end}
// }

// // 寻找左侧边界的二分查找
// func findLeftBound(nums []int, target int) int {
// 	left, right := 0, len(nums)-1 // note: [left, right]
// 	for left <= right { // note: 因为 right 是闭区间，所以可以取 =
// 		mid := left + ((right - left) >> 1) // mid = (left + right) / 2 的优化形式，防止溢出！
// 		if nums[mid] == target {
// 			right = mid - 1 // note: 收紧右侧边界以锁定左侧边界
// 		}else if nums[mid] < target {
// 			left =  mid + 1
// 		}else if nums[mid] > target {
// 			right = mid - 1
// 		}
// 	}
// 	// 返回左侧边界
// 	return left // note
// }

// // 寻找右侧边界的二分查找
// func findRightBound(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid := left + ((right - left) >> 1)
// 		if nums[mid] == target {
// 			left = mid + 1
// 		}else if nums[mid] < target {
// 			left =  mid + 1
// 		}else if nums[mid] > target {
// 			right = mid - 1
// 		}
// 	}
// 	// 返回右侧边界
// 	return right
// }