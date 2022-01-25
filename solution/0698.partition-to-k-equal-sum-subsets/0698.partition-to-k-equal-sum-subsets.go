package _698_partition_to_k_equal_sum_subsets

import (
	//"sort"
)
//
//// 方法一：n 个数是否放入桶k（n选1，放 k 次）（超时）
//func canPartitionKSubsets(nums []int, k int) bool {
//	// 特殊判断
//	n := len(nums)
//	if k > n {
//		return false
//	}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	if sum % k != 0 {
//		return false
//	}
//	target := sum / k
//	buckets := make([]int, k)
//	// 降序排列 nums，以降低时间
//	sort.Ints(nums)
//	i,j := 0, len(nums)-1
//	for ; i < j; {
//		nums[i], nums[j] = nums[j], nums[i]
//		i++
//		j--
//	}
//	return reverse(nums, buckets, target, 0)
//}
//
//func reverse(nums []int, buckets []int, target int, index int) bool {
//	if index == len(nums){
//		for _, bucket := range buckets {
//			if bucket != target {
//				return false
//			}
//		}
//		return true
//	}
//	// 穷举 nums[index] 可能装入的桶
//	for i := 0; i < len(buckets); i++ {
//		// 剪枝
//		if buckets[i] + nums[index] > target {
//			continue
//		}
//		buckets[i] += nums[index]
//		// 递归穷举下一个数字的选择
//		if reverse(nums, buckets, target, index+1) {
//			return true
//		}
//		buckets[i] -= nums[index]
//	}
//	return false
//}

// 方法二：k 个桶是否放入数n（二选一：放或不放）
func canPartitionKSubsets(nums []int, k int) bool {
	// 特殊判断
	n := len(nums)
	if k > n {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % k != 0 {
		return false
	}
	target := sum / k
	used := make([]bool, len(nums))
	// 降序排列 nums，以降低时间
	// sort.Ints(nums)
	// i,j := 0, len(nums)-1
	// for ; i < j; {
	// 	nums[i], nums[j] = nums[j], nums[i]
	// 	i++
	// 	j--
	// }
	return reverse(k, 0, 0, target, nums, used)
}

func reverse(k, bucket, start, target int, nums []int, used []bool) bool {
	if k == 0 {
		return true
	}
	if bucket == target {
		return reverse(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if nums[i] + bucket > target {
			continue
		}
		used[i] = true
		bucket += nums[i]
		// 递归穷举下一个数字是否装入当前桶
		if reverse(k, bucket, i+1, target, nums, used) {
			return true
		}
		used[i] = false
		bucket -= nums[i]
	}
	return false
}
