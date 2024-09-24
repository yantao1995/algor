package lcpr

/*
 * @lc app=leetcode.cn id=2207 lang=golang
 * @lcpr version=30204
 *
 * [2207] 字符串中最多数目的子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumSubsequenceCount(text string, pattern string) int64 {
	pre, tail := int64(0), int64(0)
	p0, p1 := int64(1), int64(1)
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[0] {
			p0++
		} else if text[i] == pattern[1] {
			pre += p0
		}
	}
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == pattern[1] {
			p1++
		} else if text[i] == pattern[0] {
			tail += p1
		}
	}
	if pattern[0] == pattern[1] {
		return max(pre, tail, p0-1, p1-1)
	}
	return max(pre, tail)
}

// @lc code=end

/*
// @lcpr case=start
// "abdcdbc"\n"ac"\n
// @lcpr case=end

// @lcpr case=start
// "aabb"\n"ab"\n
// @lcpr case=end

*/
