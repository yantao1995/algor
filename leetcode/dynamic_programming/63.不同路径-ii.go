package leetcode

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for k := range dp {
		dp[k] = make([]int, len(obstacleGrid[k]))
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				}
				if i == 0 {
					if j > 0 {
						dp[i][j] = dp[i][j-1]
					}
				} else {
					if j == 0 {
						dp[i][j] = dp[i-1][j]
					} else {
						dp[i][j] = dp[i-1][j] + dp[i][j-1]
					}
				}
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

// @lc code=end

/*
	二维dp[i][j]数组表示，从0，0 走到现在的路径总数
	因为dp数组的默认值是0，所以如果遇到四周都是障碍物的情况时，取上和左的dp值也是0
	所以只需要 在  obstacleGrid[i][j] == 0 时进行计算
*/
