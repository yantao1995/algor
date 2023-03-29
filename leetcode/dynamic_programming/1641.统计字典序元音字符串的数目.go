package leetcode

/*
 * @lc app=leetcode.cn id=1641 lang=golang
 *
 * [1641] 统计字典序元音字符串的数目
 */

// @lc code=start
func countVowelStrings(n int) int {
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, 5)
	}
	dp[0][0] = 1
	result := 1
	for i := 1; i < 5; i++ {
		dp[0][i] += dp[0][i-1]
		result += dp[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = 1
		for j := 1; j < 5; j++ {
			for m := j; m > 0; m-- {
				dp[i][j] += dp[i-1][m]
			}
			result += dp[i][j]
		}
	}
	return result
}

// @lc code=end

/*
	动态规划 dp[i][j] 表示长度为 i 以 j 结尾的字符的个数
	将 a e i o u 看做 0 1 2 3 4

	那么当前 i j 的 应该是 从 i-1长度 的 [0-j] 都可以转移过来
*/
