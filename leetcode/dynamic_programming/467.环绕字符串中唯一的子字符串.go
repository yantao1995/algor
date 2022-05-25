package leetcode

/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 */

// @lc code=start
func findSubstringInWraproundString(p string) int {
	cm := map[byte]int{}
	cm[p[0]] = 1
	dp := make([]int, len(p))
	dp[0] = 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(p); i++ {
		dp[i] = 1
		if rune(p[i])-rune(p[i-1]) == 1 ||
			rune(p[i])-rune(p[i-1]) == -25 {
			dp[i] += dp[i-1]
		}
		cm[p[i]] = max(cm[p[i]], dp[i])
	}
	result := 0
	for k := range cm {
		result += cm[k]
	}
	return result
}

// @lc code=end

/*
	dp[i]表示到当前p[i]所连续的最大长度
	因为是唯一 的 p 的 非空子串 的数量,所以只需要存到每个字符的最大连续长度
	比如 "zab"
	到z的最大长度是1
	到a的最大长度是2
	到b的最大长度是3
	所以 那么以为 z,a,b 分别作为结尾的连续字串数量就应该是 1,2,3个
	所以 如果连续，dp[i] = dp[i-1]+1
	否则 dp[i] =1
*/
