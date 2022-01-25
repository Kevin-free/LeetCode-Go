package _078_subsets

// 78. 子集：中等
// tags：数组，回溯

// 迭代法
func subsets(nums []int) [][]int {
	arr := make([][]int, 0)          // 记录结果
	arr = append(arr, []int{})       // 首先添加空集
	for i := 0; i < len(nums); i++ { // 遍历数组
		tempArr := make([][]int, 0) // 临时存储
		for _, c := range arr {     // 遍历已有子集
			temp := make([]int, 0)          // 临时存储
			temp = append(temp, c...)       // 添加子集
			temp = append(temp, nums[i])    // 添加新元素
			tempArr = append(tempArr, temp) // 添加到临时存储中
		}
		for _, c := range tempArr { // 遍历存入结果
			arr = append(arr, c)
		}
	}
	return arr
}
