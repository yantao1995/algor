package lcpr

/*
 * @lc app=leetcode.cn id=2972 lang=golang
 * @lcpr version=30204
 *
 * [2972] 统计移除递增子数组的数目 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func incremovableSubarrayCount(nums []int) int64 {
	total := 0
	pre := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			break
		}
		pre = i
	}
	if pre+1 == len(nums) {
		return int64(len(nums) * (len(nums) + 1) / 2)
	}
	for suf := len(nums) - 1; suf >= 0; suf-- {
		if suf < len(nums)-1 && nums[suf] >= nums[suf+1] {
			break
		}
		pre1 := pre
		for pre1 >= 0 && nums[pre1] < nums[suf] {
			total += pre1 + 2
			pre1--
		}
		total++
	}
	return int64(total)
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
优化思路：





Time Limit Exceeded
827/839 cases passed (N/A)

代码同 2970.统计移除递增子数组的数目 I
func incremovableSubarrayCount(nums []int) int64 {
	total := int64(0)
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
*/
