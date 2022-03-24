package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// func climbStairs(n int) int {
// 	memo := map[int]int{}
// 	var dp func(n int) int
// 	dp = func(n int) int {
// 		if n <= 2 {
// 			memo[n] = n
// 			return n
// 		}
// 		if _, ok := memo[n-2]; !ok {
// 			memo[n-2] = dp(n - 2)
// 		}
// 		if _, ok := memo[n-1]; !ok {
// 			memo[n-1] = dp(n - 1)
// 		}
// 		memo[n] = memo[n-1] + memo[n-2]
// 		return memo[n]
// 	}
// 	return dp(n)
// }

func climbStairs(n int) int {
	memo := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i <= 3 {
			memo[i] = i
		} else {
			memo[i] = memo[i-1] + memo[i-2]
		}
	}
	return memo[n]
}

// @lc code=end
