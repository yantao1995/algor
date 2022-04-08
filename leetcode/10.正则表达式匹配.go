package leetcode

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	isMatch := false
	var backtrace func(si, pi, idx int)
	backtrace = func(si, pi, idx int) {
		if isMatch {
			return
		}
		if si == len(s) && pi == len(p) {
			isMatch = true
			return
		}
		for si < len(s) && pi < len(p) {
			if s[si] == p[pi] || p[pi] == '.' {
				si++
				pi++
			} else if p[pi] == '*' {
				backtrace(si, pi+1, 0)
				if isMatch {
					return
				}
				if idx == 0 || (idx > 0 || (si > 0 && s[si] == s[si-1])) {
					si++
					backtrace(si, pi, idx+1)
					if isMatch {
						return
					}
				}
				pi++
				backtrace(si, pi, idx)
			} else if si == 0 && pi == 0 {
				pi++
			} else if si >= len(s)-1 && pi < len(p) {
				isMatch = true
				return
			} else {
				return
			}
		}
	}
	backtrace(0, 0, 0)
	return isMatch
}

// @lc code=end
