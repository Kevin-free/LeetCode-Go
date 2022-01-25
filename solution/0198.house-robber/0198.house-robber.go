package _198_house_robber


func rob(nums []int) int {
	length := len(nums)
	if nums == nil || length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	first, second := nums[0], Max(nums[0], nums[1])
	for i:=2; i<length; i++ {
		tmp := second
		second = Max(first+nums[i], second)
		first = tmp
	}
	return second

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}