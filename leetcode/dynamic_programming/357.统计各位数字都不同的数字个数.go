package leetcode

/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 统计各位数字都不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = 9
		for ii, temp := i, 9; ii > 1; ii, temp = ii-1, temp-1 {
			dp[i] *= temp
		}
		dp[i] += dp[i-1]
	}
	return dp[n]
}

// @lc code=end

/*
	动态规划
	递推：
	f(0)=1
	f(1)=9+f(0)
	f(2)=9*9+f(1)
	f(3)=9*9*8+f(2)
*/
