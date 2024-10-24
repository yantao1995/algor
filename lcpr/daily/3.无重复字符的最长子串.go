package lcpr

/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start, end := 0, 0
	countMap := map[byte]int{}
	repeatCount := 0
	for end < len(s) {
		maxLength = max(maxLength, end-start)
		countMap[s[end]]++
		if countMap[s[end]] > 1 {
			repeatCount++
		}
		for repeatCount > 0 && start <= end {
			countMap[s[start]]--
			if countMap[s[start]] == 1 {
				repeatCount--
			}
			start++
		}
		end++
	}
	maxLength = max(maxLength, end-start)
	return maxLength
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
