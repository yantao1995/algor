package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 */

// @lc code=start
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	dp := make([]int, len(startTime))
	type sep struct{ start, end, profit int }
	works := make([]sep, len(startTime))
	for i := 0; i < len(startTime); i++ {
		works[i] = sep{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}
	sort.Slice(works, func(i, j int) bool {
		return works[i].end < works[j].end
	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = works[0].profit
	for i := 1; i < len(dp); i++ {
		dp[i] = max(dp[i-1], works[i].profit)
		for j := i - 1; j >= 0; j-- {
			if works[j].end <= works[i].start {
				dp[i] = max(dp[i], dp[j]+works[i].profit)
				break
			}
		}

	}
	return dp[len(dp)-1]
}

// @lc code=end

/*
	动态规划 dp[i]表示到当前i份工作的最大收益
	按结束时间升序，然后。递推，当前天加或者不加
		当前天加 ： dp[i] = profit[i]+ dp[上一份不冲突的]
		当前天不加： dp[i] = dp[i-1]

*/
