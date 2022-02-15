package _503_last_moment_before_all_ants_fall_out_of_a_plank

// 1503. 所有蚂蚁掉下来前的最后一刻：中等
// tags：脑筋急转弯，数组，模拟
func getLastMoment(n int, left []int, right []int) int {
	max := -1
	// 向左爬的，取其最大值即可
	for _, l := range left {
		max = Max(max, l)
	}
	// 向右爬的，取 (n-蚂蚁位置）的最大值
	for _, r := range right {
		max = Max(max, n-r)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
