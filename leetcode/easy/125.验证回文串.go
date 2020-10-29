package leetcode

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start

// func isPalindrome(s string) bool {
// 	s = strings.ToLower(s)
// 	prevIndex, tailIndex := 0, len(s)-1
// 	for tailIndex > prevIndex {
// 		if !((s[tailIndex] >= 97 && s[tailIndex] <= 122) || (s[tailIndex] >= 48 && s[tailIndex] <= 57)) {
// 			tailIndex--
// 			continue
// 		}
// 		if !((s[prevIndex] >= 97 && s[prevIndex] <= 122) || (s[prevIndex] >= 48 && s[prevIndex] <= 57)) {
// 			prevIndex++
// 			continue
// 		}
// 		//	fmt.Println("_", s[prevIndex], string(s[tailIndex]))
// 		if s[prevIndex] != s[tailIndex] {
// 			return false
// 		}
// 		prevIndex++
// 		tailIndex--
// 	}
// 	return true
// }

// @lc code=end
