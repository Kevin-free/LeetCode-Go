/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
package solution
import (
	"LeetCode-Go/util"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		max = util.MaxInt(max, prices[i]-min)
		min = util.MinInt(min, prices[i])
	}
	return max
}
// @lc code=end

