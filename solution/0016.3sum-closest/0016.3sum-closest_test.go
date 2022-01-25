package _016_3sum_closest

import (
	"fmt"
	"testing"
)

func Test0016(t *testing.T) {
	nums := []int{-1,2,1,-4}
	target := 1
	res := threeSumClosest(nums,target)
	fmt.Println(res)
}