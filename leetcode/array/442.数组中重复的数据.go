package leetcode

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	result := []int{}
	n := len(nums)
	for k := range nums {
		temp := nums[k] - 1
		for temp >= n {
			temp -= n
		}
		nums[temp] += n
	}
	for k := range nums {
		if nums[k] > 2*n && nums[k] <= 3*n {
			result = append(result, k+1)
		}
	}
	return result
}

// @lc code=end
