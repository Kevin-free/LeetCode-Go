/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var res []int
	// 使用 map，map[num]index
	m := make(map[int]int, 0)
	for index, num := range nums {
		index2, exist := m[target-num]
		if exist {
			res = append(res, index)
			res = append(res, index2)
			return res
		}else {
			m[num] = index
		}
	}
	return res
	
}

// solution 1
// func twoSum(nums []int, target int) []int {
// 	var res []int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i+1; j < len(nums); j++ {
// 			if nums[i] + nums[j] == target {
// 				res = append(res, i)
// 				res = append(res, j)
// 				return res
// 			}
// 		}
// 	}
// 	return res
// }

// @lc code=end

