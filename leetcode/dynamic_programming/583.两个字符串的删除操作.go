package leetcode

/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
// func minDistance(word1 string, word2 string) int {
// 	dp := make([][]int, len(word1)+1)
// 	for k := range dp {
// 		dp[k] = make([]int, len(word2)+1)
// 	}
// 	for i := 1; i <= len(word1); i++ {
// 		dp[i][0] = i
// 	}
// 	for i := 1; i <= len(word2); i++ {
// 		dp[0][i] = i
// 	}
// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}
// 	for i := 1; i <= len(word1); i++ {
// 		for j := 1; j <= len(word2); j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else {
// 				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
// 			}
// 		}
// 	}
// 	return dp[len(word1)][len(word2)]
// }

// @lc code=end

/*
	dp[i][j] 表示当前 word1[i] 到 word2[j] 的最小删除步数
	初始化第0行，和第0列
	把 word1 和 word2 前面想象成默认加上一个 " "，这样就第0行和第0列就好初始化了
	dp[0][i] = i  , dp[j][0] = j
	因为相同的字符，转化到过去，只需要删除一个就行了。

	如果 word1[i] == word2[j] ,那么表示当前不需要删除，所以 dp[i][j] == dp[i-1][j-1]
	如果 不相等， 那么应该取  word1[i-1]到word2[j] 与 word1[i]到word2[j-1] 的最小删除数
	的基础上 再删除一个，就是当前的值，例  abc  与 ab  或者 bc 与 bce ,都只需要删除一个
*/
