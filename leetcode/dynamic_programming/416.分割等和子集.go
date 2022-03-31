package leetcode

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	total := 0
	for k := range nums {
		total += nums[k]
	}
	if total%2 == 1 || len(nums) < 2 {
		return false
	}
	half := total / 2
	can := make([]bool, half+1)
	for k := range nums {
		for j := len(can) - 1; j > 0; j-- {
			if can[j] && j+nums[k] <= half {
				can[j+nums[k]] = true
			}
		}
		if nums[k] <= half {
			can[nums[k]] = true
		}
	}
	return can[total/2]
}

// @lc code=end

/*
	两个相等的的子集，即当前数组和的一半
	数组和为奇数 或者数组长度小于2均不可能完成


*/
