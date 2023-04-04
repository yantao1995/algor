package leetcode

/*
 * @lc app=leetcode.cn id=2367 lang=golang
 *
 * [2367] 算术三元组的数目
 */

// @lc code=start
func arithmeticTriplets(nums []int, diff int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					result++
				}
			}
		}
	}
	return result
}

// @lc code=end

/*
	暴力求解即可
*/
