package leetcode

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
// func longestPalindrome(s string) int {
// 	sli := strings.Split(s, "")
// 	m := map[string]int{}
// 	for k := range sli {
// 		m[sli[k]]++
// 	}
// 	count := 0
// 	once := 0
// 	for k := range m {
// 		if m[k] == 1 {
// 			once++
// 		} else {
// 			if m[k]%2 == 1 {
// 				once++
// 			}
// 			count += (m[k] / 2)
// 		}
// 	}
// 	if once > 0 {
// 		return count*2 + 1
// 	}
// 	return count * 2
// }

// @lc code=end
