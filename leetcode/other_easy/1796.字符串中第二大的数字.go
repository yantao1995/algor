package leetcode

/*
 * @lc app=leetcode.cn id=1796 lang=golang
 *
 * [1796] 字符串中第二大的数字
 */

// @lc code=start
func secondHighest(s string) int {
	temp := make([]bool, 10)
	for i := 0; i < len(s); i++ {
		if val := s[i] - '0'; val >= 0 && val <= 9 {
			temp[val] = true
		}
	}
	count := 0
	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] {
			count++
		}
		if count == 2 {
			return i
		}
	}
	return -1
}

// @lc code=end

/*
	直接找出所有数字，然后数第二大的即可
*/
