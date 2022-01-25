package _011_container_with_most_water

import (
	"fmt"
	"testing"
)

func Test0011(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}
