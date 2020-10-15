package leetcode

/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {
	result := make([]int, len(S))
	for k := range S {
		dist := 0
		for {
			find := false
			if S[k] == C {
				find = true
			} else if k-dist >= 0 && S[k-dist] == C {
				find = true
			} else if k+dist < len(S) && S[k+dist] == C {
				find = true
			} else {
				dist++
			}
			if find {
				result[k] = dist
				break
			}
		}
	}
	return result
}

// @lc code=end
