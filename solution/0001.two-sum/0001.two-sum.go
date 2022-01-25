package _001_two_sum

func twoSum(nums []int, target int) []int {
	// length := len(nums)
	// ans := make([]int,0)
	// for i := 0; i<length; i++ {
	//     for j := i+1; j<length; j++ {
	//         if nums[i] + nums[j] == target{
	//             ans = append(ans, i, j)
	//         }
	//     }
	// }
	// return ans

	length := len(nums)
	m := make(map[int]int, length)
	result := make([]int, 0)
	for i,k := range nums{
		// 判断map中是否存在key为[target - k]的值
		if value,exist := m[target - k];exist {
			// append：尾部追加元素
			result = append(result,value)
			result = append(result, i)
			return result
		}
		m[k] = i
	}
	return result
}