package leetcode

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	count := 0
	c1, c0 := 0, 0
	for k := range strs {
		if len(strs[k]) <= m+n {
			c1, c0 = 0, 0
			for kk := range strs[k] {
				if strs[k][kk] == '1' {
					c1++
					if c1 > m {
						break
					}
				} else {
					c0++
					if c0 > n {
						break
					}
				}
			}
			if c1 <= m && c0 <= n {
				count++
			}
		}
	}
	return count
}

// @lc code=end
