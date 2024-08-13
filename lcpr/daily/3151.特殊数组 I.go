package lcpr

/*
 * @lc app=leetcode.cn id=3151 lang=golang
 * @lcpr version=30204
 *
 * [3151] 特殊数组 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isArraySpecial(nums []int) bool {
	for i := range nums {
		nums[i] &= 1
		if i > 0 && nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,1,6]\n
// @lcpr case=end

*/
