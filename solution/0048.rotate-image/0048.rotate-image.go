package _048_rotate_image

// 方法一：逐层旋转
// 将二维矩阵原地顺时针旋转 90 度
//func rotate(matrix [][]int)  {
//	length := len(matrix[0])
//	var tmp int
//	for start,end := 0, length-1; start < end;  { //层的控制
//		for s, e := start, end; s < end; { //某一层的四个点的旋转
//			tmp = matrix[start][s] //记录左上
//			matrix[start][s] = matrix[e][start] //左上=左下
//			matrix[e][start] = matrix[end][e] //左下=右下
//			matrix[end][e] = matrix[s][end] //右下=右上
//			matrix[s][end] = tmp //右上=左上
//			s++
//			e--
//		}
//		start++
//		end--
//	}
//}

// 方法二：两次旋转
// 将二维矩阵原地顺时针旋转 90 度
func rotate(matrix [][]int) {
	n := len(matrix)
	// 1. 先沿对角线翻折
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// swap(matrix[i][j], matrix[j][i]);
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. 然后翻转每一行
	for _, row := range matrix {
		reverse(row)
	}
}

// 翻转每一行
func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		// swap(nums[i], nums[j]);
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

