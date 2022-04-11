package leetcode

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	total := 0
	minLength := 0
	for i := 0; i < len(nums); i++ {
		total = 0
		for j := i; j < len(nums); j++ {
			total += nums[j]
			if total >= target {
				if j-i+1 < minLength || minLength == 0 {
					minLength = j - i + 1
				}
				break
			}
		}
	}
	return minLength
}

// @lc code=end

/*
	从当前i累加到j,遇到大于等于 target 就判断并退出
*/
