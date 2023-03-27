package leetcode

/*
 * @lc app=leetcode.cn id=1638 lang=golang
 *
 * [1638] 统计只差一个字符的子串数目
 */

// @lc code=start
func countSubstrings(s string, t string) int {
	result := 0
	isMatch := func(a, b string) bool {
		count := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
				if count > 1 {
					return false
				}
			}
		}
		return count == 1
	}
	for length := 1; length <= len(s) && length <= len(t); length++ {
		for i := 0; i+length <= len(s); i++ {
			for j := 0; j+length <= len(t); j++ {
				if isMatch(s[i:i+length], t[j:j+length]) {
					result++
				}
			}
		}
	}

	return result
}

// @lc code=end

/*
	直接暴力枚举
*/
