package leetcode

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			if nums[i] == nums[nums[i]-1] {
				continue
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for k := range nums {
		if k+1 != nums[k] {
			result = append(result, k+1)
		}
	}
	return result
}

// @lc code=end
