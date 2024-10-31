package lcpr

/*
 * @lc app=leetcode.cn id=3216 lang=golang
 * @lcpr version=30204
 *
 * [3216] 交换后字典序最小的字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getSmallestString(s string) string {
	rlt := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if (s[i]-'0')&1 == (s[i+1]-'0')&1 &&
			s[i] > s[i+1] {
			rlt[i], rlt[i+1] = rlt[i+1], rlt[i]
			break
		}
	}
	return string(rlt)
}

// @lc code=end

/*
// @lcpr case=start
// "45320"\n
// @lcpr case=end

// @lcpr case=start
// "001"\n
// @lcpr case=end

*/
