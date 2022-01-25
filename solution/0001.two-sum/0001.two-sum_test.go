package _001_two_sum

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{2, 3, 2, 1}

	fmt.Println(len(nums[:1]))
	target := 4
	res := twoSum(nums, target)
	fmt.Println(res)
}
