package easy

/*
 * @lc app=leetcode.cn id=1374 lang=golang
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */

// @lc code=start
func generateTheString(n int) string {
	result := make([]byte, n)
	result[n-1] = 'a'
	if n%2 == 0 {
		result[n-1] = 'b'
	}
	n -= 2
	for n >= 0 {
		result[n] = 'a'
		n--
	}
	return string(result)
}

// @lc code=end
