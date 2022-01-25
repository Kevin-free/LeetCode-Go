package _094_car_pooling

import (
	"fmt"
	"testing"
)

func Test1094(t *testing.T) {
	trips := [][]int{{2,1,5},{3,3,7}}
	cap := 4
	b := carPooling(trips, cap)
	fmt.Print(b)
}