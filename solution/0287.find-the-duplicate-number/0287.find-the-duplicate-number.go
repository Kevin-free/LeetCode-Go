package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	length := len(nums)
	left, right := 1, length-1
	for left < right {
	    mid := (left + right)>>1 //(left + right)/2 的位运算写法

	    cnt := 0
	    for _, num := range nums {
	        if num <= mid {
	            cnt++
	        }
	    }

	    if cnt > mid {
	        right = mid  // 缩小上界
	    }else {
	        left = mid+1 // 扩大下界
	    }
	}
	return left
}
