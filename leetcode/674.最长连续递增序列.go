package leetcode

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		for j < len(nums) {
			if nums[j] > nums[j-1] {
				if j-i+1 > max {
					max = j - i + 1
				}
				j++
				continue
			}
			break
		}
		i = j - 1
	}
	return max
}

// @lc code=end
