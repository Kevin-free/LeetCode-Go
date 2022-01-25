package _560_subarray_sum_equals_k

import "testing"

func Test0560(t *testing.T) {
	//nums := []int{1, 1, 1}
	//k := 2
	//nums := []int{1, 2, 3}
	//k := 3
	nums := []int{3, 5, 2, -2, 4, 1}
	k := 5
	res := subarraySum(nums, k)
	println(res)
}
