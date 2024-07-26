package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=2740 lang=golang
 * @lcpr version=30204
 *
 * [2740] 找出分区值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	result := nums[len(nums)-1] - nums[0]
	for i := 1; i < len(nums); i++ {
		result = min(result, nums[i]-nums[i-1])
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [100,1,10]\n
// @lcpr case=end

*/

/*
	num升序排列后
	nums1可以是升序的最大值，nums2是升序的最小值，那么可以找相邻的最小值即可
*/
