package _739_daily_temperatures

// 方法二：单调栈
// 时间复杂度：O(N) N 是温度列表的长度
// 空间复杂度：O(N) N 是温度列表的长度
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stack []int // 栈，存储下标 声明为 nil
	//stack := make([]int, n) // 栈，存储下标 声明且初始化
	for i := 0; i < n; i++ {
		// 栈不为空 && 当前天温度大于栈顶天温度
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 栈顶元素
			prevIndex := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			// 设置栈顶天的结果值
			res[prevIndex] = i - prevIndex
		}
		// 入栈
		stack = append(stack, i)
	}
	return res
}

// 方法一：暴力
// 时间复杂度：O(N*2) N 是温度列表的长度
// 空间复杂度：O(N) N 是温度列表的长度
//func dailyTemperatures(temperatures []int) []int {
//	n := len(temperatures)
//	res := make([]int, n)
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			// 如果第 j 天温度大于第 i 天温度
//			if temperatures[j] > temperatures[i] {
//				// 设置第 i 天的结果值
//				res[i] = j - i
//				break
//			}
//		}
//	}
//	return res
//}
