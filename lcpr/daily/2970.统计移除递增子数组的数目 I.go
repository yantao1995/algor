package lcpr

/*
 * @lc app=leetcode.cn id=2970 lang=golang
 * @lcpr version=30204
 *
 * [2970] 统计移除递增子数组的数目 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func incremovableSubarrayCount(nums []int) int {
	total := 0
	isIncrease := func(start, end int) bool {
		gap := 1
		last := -1
		for i := 0; i < len(nums) && gap < 2; i++ {
			if i == start {
				i = end
				continue
			}
			if nums[i] > last {
				last = nums[i]
			} else {
				gap++
			}
		}
		return gap < 2
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if isIncrease(i, j) {
				total++
			}
		}
	}
	return total
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [6,5,7,8]\n
// @lcpr case=end

// @lcpr case=start
// [8,7,6,6]\n
// @lcpr case=end

*/

/*
	直接排除递增的子数组，然后数个数就好了
*/
