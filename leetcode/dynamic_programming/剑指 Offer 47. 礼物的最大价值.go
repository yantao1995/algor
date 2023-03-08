package leetcode

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？



示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物


提示：

0 < grid.length <= 200
0 < grid[0].length <= 200

*/

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dp[i][j] = grid[i][j]
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+grid[i][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

/*
	标准动态规划， dp[i][j]的当前状态来源于 dp[i-1][j] 和 dp[i][j-1]
*/
