package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	// lb := searchRangeLB(nums, target)
	// rb := searchRangeRB(nums, target)
	// // rb := searchRangeLB(nums, target+1) - 1
	// res := make([]int, 0)
	// if len(nums) == 0 || nums[lb] != nums[rb] {
	// 	res = append(res, -1, -1)
	// } else {
	// 	res = append(res, lb, rb)
	// }
	// return res

	// 	// 目标值开始位置：为目标值的左侧边界，无此值则返回比它大的数的左侧边界
	start := searchRangeLB(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1,-1}
	}
	// 目标值结束位置：为目标值+1的左侧边界-1，无此值则返回比它大的数的左侧边界-1
	end := searchRangeLB(nums, target+1) - 1
	return []int{start, end}
}

// 在数组中寻找目标值的左侧边界
func searchRangeLB(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 在数组中寻找目标值的右侧边界
func searchRangeRB(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
