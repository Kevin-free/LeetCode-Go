package _268_missing_number

import (
	"fmt"
	"testing"
)

/**
https://leetcode-cn.com/problems/missing-number/
268. 丢失的数字
tag：数组、数学、位运算
 */

func Test0268(t *testing.T) {
	input := []int{3,0,1}
	output := missingNumber(input)
	fmt.Println(output)
}