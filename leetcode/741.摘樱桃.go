package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=741 lang=golang
 *
 * [741] 摘樱桃
 */

// @lc code=start
func cherryPickup(grid [][]int) int {
	dp := make([][]int, len(grid))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
lab1:
	for k := range grid {
		dp[k] = make([]int, len(grid[k]))
		for kk := range grid[k] {
			if grid[k][kk] == -1 {
				continue lab1
			}
			if k == kk && k == 0 {
				dp[k][kk] = grid[k][kk]
			} else if k == 0 {
				dp[k][kk] = dp[k][kk-1] + grid[k][kk]
			} else if kk == 0 {
				dp[k][kk] = dp[k-1][kk] + grid[k][kk]
			} else {
				dp[k][kk] = grid[k][kk] + max(dp[k-1][kk], dp[k][kk-1])
			}
		}
	}

	if dp[len(dp)-1][len(dp[0])-1] == 0 {
		return 0
	}
	for i, j := len(dp)-1, len(dp[0])-1; i > 0 || j > 0; {
		grid[i][j] = 0
		if j > 0 && ((dp[i][j-1]+1 == dp[i][j] && grid[i][j] == 1) || dp[i][j-1] == dp[i][j]) {
			j--
		} else {
			i--
		}
	}

	dp2 := make([][]int, len(grid))
lab2:
	for k := range grid {
		dp2[k] = make([]int, len(grid[k]))
		for kk := range grid[k] {
			if grid[k][kk] == -1 {
				continue lab2
			}
			if k == kk && k == 0 {
				dp2[k][kk] = grid[k][kk]
			} else if k == 0 {
				dp2[k][kk] = dp2[k][kk-1] + grid[k][kk]
			} else if kk == 0 {
				dp2[k][kk] = dp2[k-1][kk] + grid[k][kk]
			} else {
				dp2[k][kk] = grid[k][kk] + max(dp2[k-1][kk], dp2[k][kk-1])
			}
		}
	}

	for k := range dp {
		fmt.Println(dp[k])
	}
	fmt.Println()
	for k := range grid {
		fmt.Println(grid[k])
	}
	fmt.Println()

	for k := range dp2 {
		fmt.Println(dp2[k])
	}

	return dp[len(dp)-1][len(dp[0])-1] + dp2[len(dp2)-1][len(dp2[0])-1]
}

// @lc code=end
