package leetcode

/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	ht := map[byte]bool{}
	for k := range J {
		ht[J[k]] = true
	}
	count := 0
	for k := range S {
		if ht[S[k]] {
			count++
		}
	}
	return count
}

// @lc code=end
