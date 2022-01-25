package _303_range_sum_query_immutable

import (
	"fmt"
	"testing"
)

func Test0303(t *testing.T) {
	//numArray := NumArray{}
	numArray := Constructor([]int{1,3,5})
	ans := numArray.SumRange(0,2)
	fmt.Println(ans)

}