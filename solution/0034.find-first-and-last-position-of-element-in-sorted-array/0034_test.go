package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"fmt"
	"testing"
)

func Test0034(t *testing.T) {
	// nums := []int{5, 7, 7, 8, 8, 10}
	// target := 8
	// nums := []int{5, 7, 7, 8, 8, 10}
	// target := 6
	// nums := []int{1}
	// target := 0
	nums := []int{}
	target := 0
	res := searchRange(nums, target)
	fmt.Println("res:", res)
}
