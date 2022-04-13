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
		for m := 0; m < n; m++ {
			if dp[i][m] == 'Q' || dp[m][j] == 'Q' {
				return false
			}
			if i+m+1 < n && j+m+1 < n && dp[i+m+1][j+m+1] == 'Q' {
				return false
			}
			if i-m-1 >= 0 && j-m-1 >= 0 && dp[i-m-1][j-m-1] == 'Q' {
				return false
			}
			if i-m-1 >= 0 && j+m+1 < n && dp[i-m-1][j+m+1] == 'Q' {
				return false
			}
			if i+m+1 < n && j-m-1 >= 0 && dp[i+m+1][j-m-1] == 'Q' {
				return false
			}

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
		for j := 0; j < n; j++ {
			if canSet(i, j) {
				dp[i][j] = 'Q'
				backtrace(i+1, dp)
				dp[i][j] = '.'
			}
		}
	}
	backtrace(0, dp)
	return result
}

// @lc code=end

/*
	回溯法
	canSet 用于检测 行列斜边 是否存在 Q
	i 用于标记当前扫描的行
	j 逐一尝试当前列的每一个位置
*/
