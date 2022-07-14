package leetcode

/*
 * @lc app=leetcode.cn id=741 lang=golang
 *
 * [741] 摘樱桃
 */

// @lc code=start
func cherryPickup(grid [][]int) int {
	n, inf := len(grid), -9999
	dp := make([][]int, n+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
		for kk := range dp[k] {
			dp[k][kk] = inf
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
	dp[n-1][n-1] = grid[n-1][n-1]
	for step := 2*n - 3; step >= 0; step-- {
		for i1 := 0; i1 < n; i1++ {
			for i2 := 0; i2 < n; i2++ {
				j1, j2 = step-i1, step-i2
				if j1 >= 0 && j1 < n && j2 >= 0 && j2 < n {
					if grid[i1][j1] == -1 || grid[i2][j2] == -1 {
						dp[i1][i2] = inf
					} else {
						dp[i1][i2] = maxs(dp[i1][i2], dp[i1+1][i2], dp[i1][i2+1], dp[i1+1][i2+1]) + grid[i1][j1]
						if i1 != i2 {
							dp[i1][i2] += grid[i2][j2]
						}
					}
				}
			}
		}
	}
	if dp[0][0] < 0 {
		return 0
	}
	return dp[0][0]
}

// @lc code=end

/*
	动态规划，想象成两个人从右下角同时往左上角走,那么可以得到dp数组为dp[i1][j1][i2][j2]
	[i1,j1],[i2,j2] 分别为两个人同时走step步数时的坐标

	*因为两个人在同一时间走的步数相同，那么，两个人必然在同一条斜线上（左下到右上）
	 那么很容易得到 step = i1+j1 = i2+j2

	所以只需要i1和i2,以及step，便可以计算出纵坐标j1和j2，因此简化dp数组为 d[i1][i2]
	为了排除边界影响，定义了更宽一格来包裹dp数组

	dp[i1][i2] 表示走相同步数，此时的i1和i2在坐标轴上的横坐标

	状态转移：当前grid不为-1时，可以得到 dp[i1][i2] 来源
		如果当前的i1和i2相同，那么grid[i1][j1]和grid[i2][j2]都只能加一次
		i1,i2 同时向左  由 dp[i1][i2] 得到
		i1向左,i2向上   由 dp[i1][i2+1] 得到
		i1向上,i2向左   由 dp[i1+1][i2] 得到
		i1,i2同时向上   由 dp[i1+1][i2+1] 得到


	step :=2*n - 3
		因为：总共需要走2*n-1步,而步数是从0开始计算 (-1)，
			 并且第一步已经 由 dp[n-1][n-1] = grid[n-1][n-1] 指定了默认值 (-1)，
			 所以 初始总步数 step := 2*n-1-2


*/
