package _581_shortest_unsorted_continuous_subarray

func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	high, low, max, min := 0, length-1, nums[0], nums[length-1]
	for i:=1; i < length; i++{
		max = fmax(max, nums[i])
		min = fmin(min, nums[length-1-i])
		if nums[i] < max {
			high = i
		}
		if nums[length-1-i] > min {
			low = length-1-i
		}
	}
	if high > low {
		return high-low+1
	}
	return 0
}

func fmax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fmin(x, y int) int {
	if x < y {
		return x
	}
	return y
}