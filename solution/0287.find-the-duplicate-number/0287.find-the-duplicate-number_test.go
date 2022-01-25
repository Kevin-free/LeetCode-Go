package _287_find_the_duplicate_number

import (
	"fmt"
	"testing"
)

func Test0287(t *testing.T) {
	//input := []int{1, 3, 4, 2, 2}
	input := []int{3, 1, 3, 4, 2}
	output := findDuplicate(input)
	fmt.Println(output)

}
