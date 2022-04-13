package leetcode

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	dp := make([][]byte, n)
	for k := range dp {
		dp[k] = make([]byte, n)
		for kk := range dp[k] {
			dp[k][kk] = '.'
		}
	}
	canSet := func(i, j int) bool {
		if i == 0 {

		} else if i == n-1 {

		}

		return true
	}
	result := [][]string{}
	var backtrace func(i int, dp [][]byte)
	backtrace = func(i int, dp [][]byte) {
		if i == n {
			slice := make([]string, n)
			for k := range dp {
				slice[k] = string(dp[k])
			}
			result = append(result, slice)
			return
		}

		for ; i < n; i++ {
			for j := 0; j < n; j++ {
				if canSet(i, j) {
					dp[i][j] = 'Q'
					backtrace(i+1, dp)
					dp[i][j] = '.'
				}
			}
		}

	}
	backtrace(0, dp)
	return result
}

// @lc code=end
