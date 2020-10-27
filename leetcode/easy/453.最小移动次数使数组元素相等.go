package easy

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	maxCount := 0
	minValue := nums[0]
	for k := range nums {
		if minValue > nums[k] {
			minValue = nums[k]
		}
	}
	for k := range nums {
		maxCount += (nums[k] - minValue)
	}
	return maxCount
}

/*  计算法  case40 超时
func minMoves(nums []int) int {
	maxCount := 0
	max := -1 << 63
	max2 := -1 << 63
	maxIndex := 0
	equels := true
	for {
		equels = true
		for k := range nums {
			if k > 0 {
				if nums[k] != nums[k-1] { //max
					equels = false
				}
			}
			if max < nums[k] {
				max = nums[k]
				maxIndex = k
			}
		}
		for k := range nums { //max2
			if maxIndex != k && max2 <= nums[k] {
				max2 = nums[k]
			}
		}
		if equels {
			return maxCount
		}
		if max > max2 {
			nums[maxIndex] -= (max - max2)
			maxCount += (max - max2)
		} else {
			nums[maxIndex]--
			maxCount++
		}
		max = nums[maxIndex]
		fmt.Println("max", max, "nums", nums, "maxIndex", maxIndex, "nums[maxIndex]", nums[maxIndex])
	}
}
*/
// @lc code=end
