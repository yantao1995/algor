package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=799 lang=golang
 *
 * [799] 香槟塔
 */

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, query_row+2)
	}
	dp[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] < 1 {
				continue
			}
			dp[i+1][j] += (dp[i][j] - 1) / 2
			dp[i+1][j+1] += (dp[i][j] - 1) / 2
		}
	}
	return math.Min(1, dp[query_row][query_glass])
}

// @lc code=end

/*
	定义dp来做状态转移，
	将当前杯子里多出来的液体直接分配给下面两个杯子，然后逐层向下转移液体即可
*/
