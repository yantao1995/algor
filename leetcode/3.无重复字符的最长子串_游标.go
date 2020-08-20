package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	tailFlag := true //最后一轮游标
	cursor, start, maxLength, length := 0, 0, 0, 0
	for cursor = start; cursor < len(s); cursor++ {
		tailFlag = true
		if maxLength > cursor-start {
			tailFlag = false
		}
		maxLength = max(maxLength, cursor-start)
		for point := start; point < cursor; point++ {
			if s[cursor] == s[point] {
				tailFlag = false
				length = cursor - start
				if point == start {
					start++
				} else if cursor-point != 1 {
					start = point + 1
				} else {
					start = cursor
				}
				maxLength = max(maxLength, length)
			}
		}
	}
	if tailFlag && len(s) > 0 {
		return maxLength + 1
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
