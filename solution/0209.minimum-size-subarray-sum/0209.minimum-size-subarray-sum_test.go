package _209_minimum_size_subarray_sum

import (
	"fmt"
	"testing"
)

func Test0209(t *testing.T) {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	ans := minSubArrayLen(s, nums)
	fmt.Println(ans)

}
