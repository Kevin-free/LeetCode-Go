package _406_queue_reconstruction_by_height

import (
	"fmt"
	"testing"
)

func Test0406(t *testing.T) {
	var input = [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}

	output := reconstructQueue(input)

	fmt.Println(output) // [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
}
