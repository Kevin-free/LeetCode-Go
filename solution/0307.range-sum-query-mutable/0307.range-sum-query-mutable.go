package _307_range_sum_query_mutable

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
//func (this *NumArray) Update(i int, val int)  {
//	this.data[i] = val
//}
//
//
//func (this *NumArray) SumRange(i int, j int) int {
//	sum := 0
//	for l := i; l <= j; l++ {
//		sum += this.data[l]
//	}
//	return sum
//}


type NumArray struct {
	s_tree, nums []int
}

func Constructor(nums []int) NumArray {
	newArray := &NumArray{make([]int, len(nums) + 1), nums}
	for i, num := range nums {
		newArray.s_update(i, num)
	}
	return *newArray
}

func (this *NumArray) Update(i int, val int)  {
	this.s_update(i, val - this.nums[i])
	this.nums[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.s_sum(j + 1) - this.s_sum(i)
}

func (this *NumArray) s_update(i, val int) {
	for i, n := i + 1, len(this.nums); i <= n; i += i & -i {
		this.s_tree[i] += val
	}
}

func (this *NumArray) s_sum(i int) (res int) {
	for ; i > 0; i -= i & -i {
		res += this.s_tree[i]
	}
	return
}