package _300_longest_increasing_subsequence

import (
	"fmt"
	"testing"
)

func Test0300(t *testing.T) {
	nums := []int{10,9,2,5,3,7,101,18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}