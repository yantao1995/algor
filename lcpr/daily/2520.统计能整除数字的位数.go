package lcpr

/*
 * @lc app=leetcode.cn id=2520 lang=golang
 * @lcpr version=30102
 *
 * [2520] 统计能整除数字的位数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countDigits(num int) int {
	origin := num
	result := 0
	for num > 0 {
		s := num % 10
		if origin%s == 0 {
			result++
		}
		num /= 10
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// 7\n
// @lcpr case=end

// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// 1248\n
// @lcpr case=end

*/

/*
  求个位，然后判断是否可以除尽即可
*/
