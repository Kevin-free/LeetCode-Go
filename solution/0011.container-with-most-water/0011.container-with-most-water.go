package _011_container_with_most_water

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	res := 0
	for start < end {
		if height[start] < height[end]{
			res = max(res, height[start]*(end - start))
			start++
		}else {
			res = max(res, height[end]*(end - start))
			end--
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}
