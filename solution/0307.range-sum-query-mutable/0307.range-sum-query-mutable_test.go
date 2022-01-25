package _307_range_sum_query_mutable

import (
	"fmt"
	"testing"
)

func Test0307(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})
	res1 := numArray.SumRange(0,2)
	fmt.Println(res1)
	numArray.Update(1,2)
	res2 := numArray.SumRange(0,2)
	fmt.Println(res2)

}