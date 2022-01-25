package _075_sort_colors

func sortColors(nums []int)  {
	n := len(nums)
	i, s, e := 0, 0, n // 因为 e 设置为了 n
	for i < e { // 所以 i<e
		if nums[i] == 0 {
			nums[i], nums[s] = nums[s], nums[i]
			s++
			i++
		}else if nums[i] == 1 {
			i++
		}else {
			e-- // e为n，防止下标越界，这里要先减，在交换
			nums[i], nums[e] = nums[e], nums[i]
		}
	}
}

//func sortColors(nums []int)  {
//	n := len(nums)
//	i, s, e := 0, 0, n-1 // 也可将 e 设置为 n-1
//	for i <= e { // 所以 i可以等于e
//		if nums[i] == 0 {
//			nums[i], nums[s] = nums[s], nums[i]
//			s++
//			i++
//		}else if nums[i] == 1 {
//			i++
//		}else {
//			nums[i], nums[e] = nums[e], nums[i]
//			e-- // 因为e初值为n-1，这里就要先交换，在减
//		}
//	}
//}