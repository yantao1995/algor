package lcpr

/*
 * @lc app=leetcode.cn id=2264 lang=golang
 * @lcpr version=30204
 *
 * [2264] 字符串中最大的 3 位相同数字
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func largestGoodInteger(num string) string {
	str := ""
	for i := 0; i <= len(num)-3; i++ {
		if num[i] == num[i+1] &&
			num[i+1] == num[i+2] {
			if str == "" || str < num[i:i+3] {
				str = num[i : i+3]
			}
		}
	}
	return str
}

// @lc code=end

/*
// @lcpr case=start
// "6777133339"\n
// @lcpr case=end

// @lcpr case=start
// "2300019"\n
// @lcpr case=end

// @lcpr case=start
// "42352338"\n
// @lcpr case=end

*/
