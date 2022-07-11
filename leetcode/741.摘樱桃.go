package leetcode

import "fmt"

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
	for k := range grid {
		dp[k] = make([]int, len(grid[k]))
		for kk := range grid[k] {
			if grid[k][kk] != -1 {
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
	}

	for k := range dp {
		fmt.Println(dp[k])
	}
	fmt.Println()

	for k := len(grid) - 1; k >= 0; k-- {
		for kk := len(grid[k]) - 1; kk >= 0; kk-- {
			if grid[k][kk] != -1 {
				if k == len(grid)-1 && kk == len(grid[k])-1 {
					continue
				} else if k == len(grid)-1 {
					if dp[k][kk]+1 != dp[k][kk+1] {
						dp[k][kk] = dp[k][kk+1] + grid[k][kk]
					}
				} else if kk == len(grid[k])-1 {
					if dp[k][kk]+1 != dp[k+1][kk] {
						dp[k][kk] = dp[k+1][kk] + grid[k][kk]
					}
				} else {
					if dp[k+1][kk] > dp[k][kk+1] {
						if dp[k][kk]+1 != dp[k+1][kk] {
							dp[k][kk] = dp[k+1][kk] + grid[k][kk]
						}
					} else {
						if dp[k][kk]+1 != dp[k][kk+1] {
							dp[k][kk] = dp[k][kk+1] + grid[k][kk]
						}
					}
				}
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
	return dp[0][0]
}

// @lc code=end
