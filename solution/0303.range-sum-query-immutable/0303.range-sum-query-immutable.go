package _303_range_sum_query_immutable

//type NumArray struct {
//	data []int
//}
//
//
//func Constructor(nums []int) NumArray {
//	return NumArray{nums}
//}
//
//
//func (this *NumArray) SumRange(i int, j int) int {
//	sum := 0
//	for k:= i; k <=j; k++ {
//		sum += this.data[k]
//	}
//	return sum
//}

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	a := NumArray{}
	a.sum = append(a.sum, 0)
	for i := 1; i <= len(nums); i++ {
		a.sum = append(a.sum, a.sum[i-1]+nums[i-1])
	}
	return a
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
