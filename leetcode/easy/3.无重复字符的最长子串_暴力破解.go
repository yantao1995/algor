package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
// func lengthOfLongestSubstring(s string) int {
// 	length := 0
// 	for i := 0; i < len(s); i++ {
// 		currentlength := 1
// 		sli := append([]byte{}, byte(s[i]))
// 		for j := i + 1; j < len(s); j++ {
// 			for _, v := range sli {
// 				if byte(v) == s[j] {
// 					goto comehere
// 				}
// 			}
// 			currentlength++
// 			sli = append(sli, byte(s[j]))
// 		}
// 	comehere:
// 		if length < currentlength {
// 			length = currentlength
// 		}
// 	}
// 	return length
// }

// @lc code=end
