package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
func bestTeamScore(scores []int, ages []int) int {
	type sa struct {
		score, age int
	}
	sas := make([]sa, len(scores))
	for k := range scores {
		sas[k] = sa{
			score: scores[k],
			age:   ages[k],
		}
	}
	sort.Slice(sas, func(i, j int) bool {
		return sas[i].age < sas[j].age ||
			(sas[i].age == sas[j].age && sas[i].score < sas[j].score)
	})

	dp := make([]int, len(sas))
	dp[0] = sas[0].score
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := sas[0].score
	for i := 1; i < len(dp); i++ {
		dp[i] = sas[i].score
		for j := i - 1; j >= 0; j-- {
			if sas[i].age > sas[j].age && sas[j].score > sas[i].score {
				continue
			}
			dp[i] = max(dp[i], sas[i].score+dp[j])
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

// @lc code=end

/*
	dp[i] 当前球员与前面所有球员组成无矛盾球队的最高得分
	按年龄从小到大,分数从小到大升序
	然后动态规划来选择，当前球员与前面无矛盾球员所能得到的最高分数
	用result记录最高得分即可
*/
