package leetcode

/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	count := 0
	cts, currentCount, prevSymbol := []int{}, 0, s[0]
	for i := 0; i < len(s); i++ {
		if s[i] == prevSymbol {
			currentCount++
		} else {
			cts = append(cts, currentCount)
			prevSymbol = s[i]
			currentCount = 1
		}
	}
	cts = append(cts, currentCount)
	for i := 1; i < len(cts); i++ {
		if cts[i] < cts[i-1] {
			count += cts[i]
		} else {
			count += cts[i-1]
		}
	}
	return count
}

//Time Limit Exceeded
//90/90 cases passed (N/A)
// func countBinarySubstrings(s string) int {
// 	count := 0
// 	for i := 0; i < len(s); i++ {
// 		iCount := 1
// 		JCount := 0
// 		for j := i + 1; j < len(s); j++ {
// 			if JCount != 0 && s[j] == s[i] {
// 				break
// 			} else {
// 				if s[j] == s[i] {
// 					iCount++
// 				} else {
// 					JCount++
// 				}
// 				if iCount == JCount {
// 					count += iCount
// 					i += iCount - 1		//优化了计数次数
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// @lc code=end
