package _698_partition_to_k_equal_sum_subsets

import (
	"fmt"
	"testing"
)

func Test0698(t *testing.T) {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	res := canPartitionKSubsets(nums, k)
	fmt.Println(res)
}