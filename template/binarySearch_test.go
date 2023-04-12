package template

import (
	"fmt"
	"testing"
)

func TestBinarySearchWithR(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 9}
	target := 2
	res := binarySearchWithR(nums, target)
	fmt.Printf("res: %v\n", res)
	// push
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 9}
	target := 2
	res := BinarySearch(nums, target)
	fmt.Printf("res: %v\n", res)
	// push
}
