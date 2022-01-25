package _136_single_number

func singleNumber(nums []int) int {
	res := nums[0]
	for _, v := range nums {
		res = res ^ v
	}
	return res
}