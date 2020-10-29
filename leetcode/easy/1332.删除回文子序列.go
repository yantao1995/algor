package leetcode

/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] 删除回文子序列
 */

// @lc code=start
//回文子序列
func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	countA, countB := 0, 0
	if checkPalindrome1332(s) {
		return 1
	}
	for k := range s {
		if s[k] == 'a' && countA == 0 {
			countA = 1
		}
		if s[k] == 'b' && countB == 0 {
			countB = 1
		}
		if countA == 1 && countB == 1 {
			return countA + countB
		}
	}
	return countA + countB
}
func checkPalindrome1332(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 回文子串
// Testcase
// "bbaabaaa"
// Answer
// 3
// Expected Answer
// 2
// func removePalindromeSub(s string) int {
// 	count := 0
// 	for s != "" {
// 		length := len(s)
// 		for length > 0 {
// 			for i := 0; i+length <= len(s); i++ {
// 				if checkPalindrome1332(s[i : i+length]) {
// 					s = s[:i] + s[i+length:]
// 					goto come
// 				}
// 			}
// 			length--
// 		}
// 	come:
// 		count++
// 	}
// 	return count
// }

// @lc code=end
