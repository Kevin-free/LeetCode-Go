package template

// 二分查找目标值的下标
// 带右边界的普通二分查找
func binarySearchWithR(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // note: [left, right]

	if nums[left] > target || nums[right] < target {
		return -1
	}
	
	for left <= right { // note: <=
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := left + ((right - left) >> 1)
	if nums[left] > target || nums[right] < target {
		return -1
	}
	for left < right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
