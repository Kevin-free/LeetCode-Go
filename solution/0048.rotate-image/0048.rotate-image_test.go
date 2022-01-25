package _048_rotate_image

import (
	"fmt"
	"testing"
)

func Test0048(t *testing.T) {

	nums := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	rotate(nums)
	fmt.Println(nums)
}