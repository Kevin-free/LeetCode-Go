package _300_longest_increasing_subsequence

// 300. 最长递增子序列：中等
// tags：数组、动态规划、二分查找

// DP
//- 时间复杂度：O(N^2)
//- 空间复杂度：O(N)
//func lengthOfLIS(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	dp := make([]int, n)
//	dp[0] = 1
//	maxans := 1
//	for i := 1; i < n; i++ {
//		// 初始化
//		dp[i] = 1
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				// 状态转移
//				dp[i] = util.MaxInt(dp[i], dp[j]+1)
//			}
//		}
//		maxans = util.MaxInt(maxans, dp[i])
//	}
//	return maxans
//}

// DP + 贪心 + 二分
//- 时间复杂度：O(NlogN)
//- 空间复杂度：O(N)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	tail := make([]int, n)
	// end 表示有序数组 tail 的最后一个已经赋值元素的索引
	end := 0
	// 初始化
	tail[end] = nums[0]
	for i := 1; i < n; i++ {
		// 加入 tail 数组末尾，并更新 end = end + 1
		if nums[i] > tail[end] {
			//tail = append(tail, nums[i]) // fixed：不能用 append ，
			//直接添加在那个元素的后面，所以 end 先加 1
			end++
			tail[end] = nums[i]
		}else {
			// 在 `tail` 数组中二分查找比 `nums[i]` 大的数的左侧边界 `k`，并更新 `tail[k]=nums[i]`
			left, right := 0, end-1
			for left <= right {
				mid := left + ((right - left)>>1)
				if tail[mid] < nums[i] {
					left = mid + 1
				}else {
					right = mid - 1
				}
			}
			tail[left] = nums[i]
		}
	}
	// 此时 end 是有序数组 tail 最后一个元素的索引
	// 题目要求返回的是长度，因此 +1 后返回
	end++
	return end
}
