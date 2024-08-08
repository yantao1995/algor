package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=2974 lang=golang
 * @lcpr version=30204
 *
 * [2974] 最小数字游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,5]\n
// @lcpr case=end

*/

/*
	读题就完事了
*/
