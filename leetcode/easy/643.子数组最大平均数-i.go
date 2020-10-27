package easy

/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	total := 0.0
	for i := 0; i < k; i++ {
		total += float64(nums[i])
	}
	maxValue := total / float64(k)
	for i := k; i < len(nums); i++ {
		total = total - float64(nums[i-k]) + float64(nums[i])
		if total/float64(k) > maxValue {
			maxValue = total / float64(k)
		}
	}
	return maxValue
}

// @lc code=end
