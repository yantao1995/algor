package leetcode

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for k := range dp {
		dp[k] = make([]int, len(word2)+1)
		dp[k][0] = k
	}
	for k := range dp[0] {
		dp[0][k] = k
	}
	min := func(val int, nums ...int) int {
		for i := 0; i < len(nums); i++ {
			if val > nums[i] {
				val = nums[i]
			}
		}
		return val
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// @lc code=end

/*
	参考官方题解  动态规划
	dp[i][j] 代表 word1 中前 i 个字符，变换到 word2 中前 j 个字符，最短需要操作的次数
	需要考虑 word1 或 word2 一个字母都没有，即全增加/删除的情况，所以预留 dp[0][j] 和 dp[i][0]

	状态转移
	增，dp[i][j] = dp[i][j - 1] + 1
	删，dp[i][j] = dp[i - 1][j] + 1
	改，dp[i][j] = dp[i - 1][j - 1] + 1
	按顺序计算，当计算 dp[i][j] 时，dp[i - 1][j] ， dp[i][j - 1] ， dp[i - 1][j - 1] 均已经确定了
	配合增删改这三种操作，需要对应的 dp 把操作次数加一，取三种的最小
	如果刚好这两个字母相同 word1[i - 1] = word2[j - 1] ，那么可以直接参考 dp[i - 1][j - 1] ，操作不用加一
*/
