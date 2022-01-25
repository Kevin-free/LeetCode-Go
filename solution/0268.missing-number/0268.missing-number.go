package _268_missing_number

/**
https://leetcode-cn.com/problems/missing-number/
268. 丢失的数字
tag：数组、数学、位运算
 */

// 解法一：数学
//func missingNumber(nums []int) int {
//	n := len(nums)
//	expectedSum := n*(n+1)/2 // 求和公式
//	actualSum := 0
//	for _,num := range nums {
//		actualSum += num
//	}
//	res := expectedSum - actualSum
//	return res
//}

// 解法二：位运算
func missingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res ^= num
		res ^= i
	}
	return res
}