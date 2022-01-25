package _096_unique_binary_search_trees

func numTrees(n int) int {
	return gn(n)
}

func gn(n int) int{
	if n==0{
		return 1
	}
	if n==1{
		return 1
	}
	var res int
	for i:=0; i <= n-1; i++{
		res += gn(i)*gn(n-1-i)
	}
	return res
}
