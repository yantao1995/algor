package leetcode

/*
 * @lc app=leetcode.cn id=1780 lang=golang
 *
 * [1780] 判断一个数字是否可以表示成三的幂的和
 */

// @lc code=start
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

// @lc code=end

/*
	转换成3进制，只包含 0和1
	即： 3^0+3^2+3^4 = 10101
*/
