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

/*
	dp数组memorandum记录字符串是否为回文子串。
	end 向后遍历，start向前扩张，只需要 s[start] =s[end]，再判断start和end中间的
	字符串是否为回文字符串，就可以知道当前 s[start:end]是否为回文字符串。
*/
