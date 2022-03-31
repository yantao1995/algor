package leetcode

/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	flag := make([]bool, len(nums))
	max := 0
	var dp func(count, current int)
	dp = func(count, current int) {
		if count == len(nums) {
			if current > max {
				max = current
			}
			return
		}
		for i := range flag {
			if !flag[i] {
				temp := nums[i]
				for l := i - 1; l >= 0; l-- {
					if !flag[l] {
						temp *= nums[l]
						break
					}
				}
				for r := i + 1; r < len(flag); r++ {
					if !flag[r] {
						temp *= nums[r]
						break
					}
				}
				flag[i] = true
				dp(count+1, current+temp)
				flag[i] = false
			}
		}
	}
	dp(0, 0)
	return max
}

// @lc code=end

/*


 */
