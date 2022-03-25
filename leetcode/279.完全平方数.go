package leetcode

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	memo := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		memo[i*i] = 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var bfs func(n int) int
	bfs = func(n int) int {
		if n < 4 {
			memo[n] = n
			return n
		}
		if memo[n] > 0 {
			return memo[n]
		}
		m := n
		for i := 1; i <= n/2; i++ {
			m = min(m, bfs(n-i)+bfs(i))
		}
		memo[n] = m
		return m
	}
	bfs(n)
	return memo[n]
}

// @lc code=end
// 1 4 9 16 25 36 49
