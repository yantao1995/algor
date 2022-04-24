package leetcode

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return len(s)
	}
	dp := make([][]int, len(s))
	for k := range dp {
		dp[k] = make([]int, len(t))
	}
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(s); i++ {
		dp[i][0] = dp[i-1][0]
		if s[i] == t[0] {
			dp[i][0]++
		}
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t) && j <= i; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)-1][len(t)-1]
}

// @lc code=end

/*
	dp[i][j] 表示 s[i] 到 t[j] 的出现的个数
	初始化第0行，因为 s只有一个字符的时候，只能匹配 t[0] 所以只需要初始化dp[0][0]
	初始化第0列，因为 t只有一个字符的时候，只要s[i] == t[0]，就能匹配 +1

	例子：
			b  a  g
		b   1  0  0
		a   1  1  0
		b   2  1  0
		g   2  1  1
		b   3  1  1
		a   3  4  1
		g   3  4  5

	如果 s[i] != t[j] ,那么当前状态转移为 dp[i][j] = dp[i-1][j] ，
		因为当前匹配结果应该和 s[i-1] 与 t[j] 的匹配结果一样，因为当前结束字符不一样，无法匹配
	如果 s[i] == t[j] ,那么当前状态转移为 dp[i][j] = dp[i-1][j-1] + dp[i-1][j] ，
		因为与上面一样，包含 s[i-1] 与 t[j-1] 的匹配结果，也包含 s[i-1] 与 t[j] 的匹配结果
		而 s[i]与t[j-1]则不是前置状态，因为t是全匹配，不能依赖t[j-1]的状态
*/
