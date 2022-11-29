package leetcode

/*
 * @lc app=leetcode.cn id=1758 lang=golang
 *
 * [1758] 生成交替二进制字符串的最少操作数
 */

// @lc code=start
func minOperations(s string) int {
	count1, count2 := 0, 0
	s2 := s
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] != '1' {
			count1++
		}
		if i%2 == 1 && s[i] != '0' {
			count1++
		}

		if i%2 == 0 && s2[i] != '0' {
			count2++
		}
		if i%2 == 1 && s2[i] != '1' {
			count2++
		}
	}
	if count1 < count2 {
		return count1
	}
	return count2
}

// @lc code=end

/*
	分别统计以1开头和以0开头的需要反转的个数
*/
