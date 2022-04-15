package leetcode

/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	isIncrease := true //是否增加
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[0])
		} else {
			if nums[i] == nums[i-1] {
				continue
			}
			if (isIncrease && nums[i] > result[len(result)-1]) ||
				(!isIncrease && nums[i] < result[len(result)-1]) {
				if len(result) == 1 {
					result = append(result, nums[i])
				} else {
					result[len(result)-1] = nums[i]
				}
			} else {
				result = append(result, nums[i])
				isIncrease = !isIncrease
			}
		}
	}
	return len(result)
}

// @lc code=end

/*
	每次都贪当前递增/递减序列的最大/小值
	因为需要严格保证 + -顺序，所以
	在 + 的时候，取最大值，增序列，就取序列最大值
	在 - 的时候，取最小值，减序列，就取序列的最小值
	这样就能最大程度的保障下一个序列能够满足
	注意第一个元素的选取
*/
