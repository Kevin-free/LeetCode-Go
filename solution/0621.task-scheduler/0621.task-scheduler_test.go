package _621_task_scheduler

import (
	"fmt"
	"testing"
)

func Test0621(t *testing.T) {
	input := []byte{'A', 'A', 'B', 'B'}
	//input := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D'}
	output := leastInterval(input, 0)
	fmt.Println(output)
}
