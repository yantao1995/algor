package lcpr

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=3131 lang=golang
 * @lcpr version=30204
 *
 * [3131] 找出与数组相加的整数 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func addedInteger(nums1 []int, nums2 []int) int {
	return slices.Min(nums2) - slices.Min(nums1)
}

// @lc code=end

/*
// @lcpr case=start
// [2,6,4]\n[9,7,5]\n
// @lcpr case=end

// @lcpr case=start
// [10]\n[5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n[1,1,1,1]\n
// @lcpr case=end

*/

/*
	找最小值减即可
*/
