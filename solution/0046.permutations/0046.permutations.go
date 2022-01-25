package _046_permutations
//
//import "fmt"
//
//// 46. 全排列：中等
//// tags：数组，回溯
//
//// 记录返回结果
////var res [][]int
//
//// 输入一组不重复的数，返回它们的全排列
//func permute(nums []int) [][]int {
//	length := len(nums)
//	// 记录返回结果
//	res := make([][]int, 0)
//
//	// note：初始化 res
//	//res = [][]int{}
//	if length == 0 {
//		return res
//	}
//
//	// 记录数是否使用过
//	used := make([]bool, length)
//	// 记录路径
//	path := make([]int, 0)
//
//	dfs(nums, 0, path, used, res)
//	fmt.Println("res:",res)
//	fmt.Println("used:",used)
//	fmt.Println("path:",path)
//	return res
//}
//
//// dfs 遍历
//func dfs(nums []int, depth int, path []int, used []bool, res [][]int) {
//	// 结束条件
//	if depth == len(nums) {
//		// note：因为 path 是地址引用，防止 path 的改变影响 res，所以需要 copy 到 tmp
//		tmp := make([]int, len(path))
//		copy(tmp, path)
//		res = append(res, tmp)
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		// 排除不合法的选择
//		if !used[i] {
//			// 做选择
//			path = append(path, nums[i])
//			used[i] = true
//			// 进入下一层决策树
//			dfs(nums, depth+1, path, used, res)
//			// 回退选择
//			used[i] = false
//			path = path[:len(path)-1]
//		}
//	}
//}
