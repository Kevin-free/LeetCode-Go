package _015_3sum

import (
	"fmt"
	"testing"
)

func Test0015(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
