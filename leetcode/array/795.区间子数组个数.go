package leetcode

/*
 * @lc app=leetcode.cn id=795 lang=golang
 *
 * [795] 区间子数组个数
 */

// @lc code=start
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	result := 0
	il, ir := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > right {
			il = i
		}
		if nums[i] >= left {
			ir = i
		}
		result += ir - il
	}
	return result
}

// @lc code=end

/*
	枚举右端点:以当前索引i结尾的数组的个数
	ir 以索引ir结尾的数组
	il 以索引il+1为起始的数组
*/
