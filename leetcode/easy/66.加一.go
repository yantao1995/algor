package easy

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	bits := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+bits >= 10 {
			digits[i] = (digits[i] + bits) % 10
			bits = 1
			if i == 0 {
				digits = append([]int{bits}, digits...)
			}
		} else {
			digits[i] = digits[i] + bits
			return digits
		}
	}
	return digits
}

// @lc code=end
