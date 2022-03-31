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
		memo[n] = n
		for i := 1; i*i <= n; i++ {
			memo[n] = min(memo[n], bfs(n-i*i)+1)
		}
		return memo[n]
	}
	bfs(n)
	return memo[n]
}

// @lc code=end
// 1 4 9 16 25 36 49

/*
	先初始化 memo 数组，直接给当前所有完全平方数所需数量初始化为1
	然后使用bfs，
	在当前值n和当前值n减去一个完全平方数的个数中取最小值。
	即 memo[n] = min(memo[n], bfs(n-i*i)+1)
	+1是因为 n-i*i 表示 选取当前 i*i 的值作为完全平方数
	同时记录 memo[减去之后的n] 的最小值，作为备忘录
*/
