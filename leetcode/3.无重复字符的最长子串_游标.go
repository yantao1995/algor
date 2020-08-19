package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	cursor, start, maxLength, length := 0, 0, 0, 0
	for cursor = start; cursor < len(s); cursor++ {
		fmt.Println(cursor, start)
		maxLength = max(maxLength, cursor-start)
		for point := start; point < cursor; point++ {
			if s[cursor] == s[point] {
				length = max(cursor-point, point-start)
				if cursor-point == maxLength {
					fmt.Println("aaaa", point)
					start = point - 1
				}
				maxLength = max(length, maxLength)
				start = cursor
				break
			}
		}
	}
	maxLength = max(maxLength, cursor-start)
	fmt.Println(cursor, start)
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
