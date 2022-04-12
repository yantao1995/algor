package leetcode

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	m := map[int]bool{} //
	calc := func(n int) int {
		result := 0
		for n > 0 {
			n1 := n % 10
			n /= 10
			result += n1 * n1
		}
		return result
	}
	for !m[n] {
		m[n] = true
		n = calc(n)
		if n == 1 {
			return true
		}
	}
	return false
}

// @lc code=end

/*
	因为会出现无限循环的情况，所以用m来存过去已经出现过的数
	calc用来计算下一轮的结果
*/
