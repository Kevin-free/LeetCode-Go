package _509_fibonacci_number

// 方法一：递归
//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 || n == 2 {
//		return 1
//	}
//	return fib(n-1) + fib(n-2)
//}

// 方法二：DP
//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//	dp := make([]int, n+1)
//	dp[0], dp[1] = 0, 1
//	for i := 2; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[n]
//}

// 状压DP
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	pre := 1
	cur := 1
	for i := 3; i <= n; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
