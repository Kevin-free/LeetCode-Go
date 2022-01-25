package _560_subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	n := len(nums)
	// map：前缀和 -> 该前缀和出现的次数
	preSum := make(map[int]int, 0)
	// base case
	preSum[0] = 1
	res, sum0_i := 0, 0
	for i := 0; i < n; i++ {
		sum0_i += nums[i]
		// 这是我们想找的前缀和 nums[0..j]
		sum0_j := sum0_i - k
		// 如果前面有这个前缀和，则直接更新答案
		if value, exist := preSum[sum0_j]; exist {
			res += value
		}
		// 把前缀和 nums[0..i] 加入并记录出现次数
		preSum[sum0_i] = preSum[sum0_i] + 1
	}
	return res
}

//func subarraySum(nums []int, k int) int {
//    n := len(nums)
//	// 构造前缀和
//    preSum := make([]int, n+1)
//    preSum[0] = 0
//    for i := 0; i<n; i++ {
//        preSum[i+1] = preSum[i] + nums[i]
//    }
//
//    var res int
//	// 穷举所有子数组
//	// 因为前缀和数组长度 n+1，所以范围为 [1,n]
//    for i := 1; i <= n; i++ {
//        for j := 0; j < i; j++ {
//			// 子数组 nums[j..i-1] 的元素和
//            if preSum[i] - preSum[j] == k {
//                res++
//            }
//        }
//    }
//    return res
//}