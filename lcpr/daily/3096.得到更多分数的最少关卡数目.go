package lcpr

/*
 * @lc app=leetcode.cn id=3096 lang=golang
 * @lcpr version=30204
 *
 * [3096] 得到更多分数的最少关卡数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumLevels(possible []int) int {
	bob := 0
	for k := range possible {
		if possible[k] == 0 {
			possible[k] = -1
		}
		bob += possible[k]
	}
	alice := 0
	for i := 0; i < len(possible)-1; i++ {
		alice += possible[i]
		bob -= possible[i]
		if alice > bob {
			return i + 1
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0]\n
// @lcpr case=end

*/
