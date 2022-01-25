package _094_car_pooling
//
//func carPooling(trips [][]int, capacity int) bool {
//	// 最多有 1000 个车站
//	nums := make([]int, 1001)
//	// 构造差分数组
//	initial(nums)
//	// 获取每条计划
//	for _, trip := range trips {
//		// 乘客数
//		val := trip[0]
//		// 第 trip[1] 站乘客上车
//		i := trip[1]
//		// 第 trip[2] 站乘客已经下车，
//		// 即乘客在车上的区间是 [trip[1], trip[2] - 1]
//		j := trip[2] - 1
//		// 进行区间操作
//		increment(i, j, val)
//	}
//	// 获取结果数组
//	res := result()
//	// 客车自始至终都不应该超载
//	for _, val := range res {
//		if val > capacity {
//			return false
//		}
//	}
//	return true
//}
//
//// 差分数组工具类
////type Difference struct {
////	// 差分数组
////	diff []int
////}
//
//// 差分数组
//var diff []int
//
//// 输入一个初始数组，区间操作将在这个数组上进行
//func initial(nums []int) {
//	if len(nums) <= 0 {
//		return
//	}
//	diff = make([]int, len(nums))
//	// 根据初始数组构造差分数组
//	diff[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		diff[i] = nums[i] - nums[i-1]
//	}
//}
//
//// 给闭区间 [i,j] 增加 val （可以是负数）
//func increment(i, j, val int) {
//	diff[i] += val
//	if j+1 < len(diff) {
//		diff[j+1] -= val
//	}
//}
//
//// 返回结果数组
//func result() []int {
//	res := make([]int, len(diff))
//	// 根据差分数组构造结果数组
//	res[0] = diff[0]
//	for i := 1; i < len(diff); i++ {
//		res[i] = res[i-1] + diff[i]
//	}
//	return res
//}
