package easy

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	total := [][]int{}
	backTrack(nums, 0, &total)
	return total
}

func backTrack(nums []int, i int, total *[][]int) {
	if i == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*total = append(*total, temp)
	}
	for k := i; k < len(nums); k++ {
		nums[i], nums[k] = nums[k], nums[i]
		backTrack(nums, i+1, total)
		nums[i], nums[k] = nums[k], nums[i]
	}
}

// @lc code=end
