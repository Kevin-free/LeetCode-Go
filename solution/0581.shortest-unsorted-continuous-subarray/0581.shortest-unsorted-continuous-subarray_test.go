package _581_shortest_unsorted_continuous_subarray

import "testing"

func Test0581(t *testing.T) {
	nums := []int{5,4,3,2,1}
	res := findUnsortedSubarray(nums)
	println(res)
}