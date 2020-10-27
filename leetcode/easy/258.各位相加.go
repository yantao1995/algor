package easy

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	weight, temp := 1, 0
	for num >= 10 {
		temp, weight = 0, 1
		for weight <= num {
			weight *= 10
		}
		weight /= 10
		for weight >= 1 {
			temp += (num / weight)
			num %= weight
			weight /= 10
		}
		num = temp
	}
	return num
}

// @lc code=end
