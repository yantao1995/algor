package leetcode

/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	val := append([]int{1}, append(nums, 1)...)
	dp := make([][]int, len(val))
	for k := range dp {
		dp[k] = make([]int, len(val))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := len(val) - 1; i >= 0; i-- {
		for j := i + 1; j < len(val); j++ {
			for mid := i + 1; mid < j; mid++ {
				dp[i][j] = max(dp[i][j], dp[i][mid]+dp[mid][j]+val[i]*val[j]*val[mid])
			}
		}
	}
	return dp[0][len(dp[0])-1]
}

// @lc code=end

/*
	//自底向上 dp右下角向上 dp[i][j]依赖 dp[i][mid] dp[mid][j] 所以先算后面的
	dp[i][j] 表示 i到j 区间内的最大值

	之所以 i 要从最后一个开始，是因为 dp[i][mid]+dp[mid][j]  mid是介于 i与j之间
	而且 必须要先知道 dp[mid][j]的最大值，才能推导出 dp[i][j]
	对应 i,j 矩阵，就是从右下角推到左上角

*/
