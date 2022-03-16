package leetcode

/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	memorandum := map[string]bool{}
	count := 0
	for end := range s {
		for start := end; start >= 0; start-- {
			if end == start ||
				(s[end] == s[start] && (end-start == 1 || memorandum[string(s[start+1:end])])) {
				count++
				memorandum[string(s[start:end+1])] = true
			}
		}
	}
	return count
}

// @lc code=end
