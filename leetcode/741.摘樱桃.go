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
	n := len(grid)
	dp := make([][][]int, 2*n)
	for k := range dp {
		dp[k] = make([][]int, n)
		for kk := range dp[k] {
			dp[k][kk] = make([]int, n)
		}
	}
	maxs := func(a int, b ...int) int {
		for k := range b {
			if b[k] > a {
				a = b[k]
			}
		}
		return a
	}
	j1, j2 := 0, 0
	dp[0][0][0] = grid[0][0]
	for step := 1; step < 2*n; step++ {
		for i1 := 0; i1 < n; i1++ {
			for i2 := 0; i2 < n; i2++ {
				if j1, j2 = step-i1, step-i2; j1 >= 0 && j1 < n && j2 >= 0 && j2 < n {
					if grid[i1][j1] == -1 || grid[i2][j2] == -1 {
						dp[step][i1][i2] = -9999
					} else {
						t1, t2, t3, t4 := -9999, -9999, -9999, -9999
						if i1 > 0 {
							t1 = dp[step-1][i1-1][i2]
						}
						if i2 > 0 {
							t2 = dp[step-1][i1][i2-1]
						}
						if i1 > 0 && i2 > 0 {
							t3 = dp[step-1][i1-1][i2-1]
						}
						t4 = dp[step-1][i1][i2]
						t1 = maxs(t1, t2, t3, t4) + grid[i1][j1]
						if i1 != i2 {
							t1 += grid[i2][j2]
						}
						dp[step][i1][i2] = maxs(dp[step][i1][i2], t1)
					}
				}
			}
		}
	}
	for k := range dp {
		for kk := range dp[k] {
			fmt.Println(dp[k][kk])
		}
		fmt.Println("------------------------")
	}
	if dp[2*n-1][n-1][n-1] < 0 {
		return 0
	}
	return dp[2*n-1][n-1][n-1]
}

// @lc code=end
