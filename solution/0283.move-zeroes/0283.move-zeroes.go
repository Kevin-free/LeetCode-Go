package _283_move_zeroes

func moveZeroes(nums []int)  {
	//if nums == nil {
	//	return
	//}
	//b := 0
	//for a := 0; a < len(nums); a++ {
	//	if nums[a] != 0 {
	//		nums[b] = nums[a]
	//		b++
	//	}
	//}
	//for i := b; i < len(nums); i++ {
	//	nums[i] = 0
	//}


	if nums == nil {
		return
	}
	b := 0
	for a:=0; a< len(nums); a++ {
		if nums[a] != 0 {
			nums[a], nums[b] = nums[b], nums[a]
			b++
		}
	}
}