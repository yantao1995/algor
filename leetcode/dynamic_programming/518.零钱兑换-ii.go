package leetcode

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)) // dp[i][j]
	for k := range dp {
		dp[k] = make([]int, amount+1)
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if i == 0 {
				if j%coins[i] == 0 {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = dp[i-1][j]
				if j >= coins[i] {
					dp[i][j] += dp[i][j-coins[i]]
				}
			}
		}
	}
	return dp[len(coins)-1][amount]
}

// @lc code=end

/*
	dp[i][j] 表示 有 0-i 种硬币，可以组成 amount为 j 的组合数量
	初始化第0行数组，即 当前只有 coins[0] 这一种硬币，可以分别组成 [0,amount]的种类
	因为只有一种硬币，所以只需要 取余为0，那么就可以组成 1种

	j表示当前的amount
	i 增加1 ，表示新增了一种硬币，那么 当前硬币，用还是不用，是一种选择
	如果j都小于coins[i]，那么硬币就用不上，直接就取 dp[i-1][j]，就直接用上一种硬币
	如果j够用，那么就可以增加一枚 coins[i] ,那么就相当于 不用的种数  + 用的种数
	由于不用的种数已经被赋值给了， dp[i][j] ，所以只需要加上 dp[i][j-coins[i]]
	 ::当前j可以由  j-coins[i] 加  coins[i]得到，也可以由 dp[i-1][j]得到
	dp[i][j] += dp[i][j-coins[i]]
*/
