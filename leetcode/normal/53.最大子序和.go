package leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	maxLength := nums[0]
	currentTotal := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && currentTotal+nums[i] < 0 {
			currentTotal = -nums[i]
			maxLength = max(maxLength, nums[i])
		}
		currentTotal += nums[i]
		if nums[i] >= 0 {
			maxLength = max(maxLength, currentTotal)
		}
	}
	return maxLength
}

// @lc code=end

/*
再此基础上优化执行逻辑


因为是连续的，所以一直向后累加，当前连续数组的和如果小于0，
说明会拖累后面的数组，直接抛弃当前的数组。然后从当前位置作为起点再次向后处理
func maxSubArray(nums []int) int {
	maxLength := nums[0]
	currentTotal := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if currentTotal+nums[i] < 0 {
				currentTotal = 0
				maxLength = max(maxLength, nums[i])
			} else {
				currentTotal += nums[i]
			}
		} else {
			currentTotal += nums[i]
			maxLength = max(maxLength, currentTotal)
		}

	}
	return maxLength
}

超时
计算累积到当前nums的和，然后与前面的值做差得到连续子数组和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	maxLength := nums[0]
	for i := 0; i < len(nums); i++ {
		maxLength = max(maxLength, nums[i])
		for j := i + 1; j < len(nums); j++ {
			maxLength = max(maxLength, nums[j]-nums[i])
		}
	}
	return maxLength
}
*/
