package easy

/*
 * @lc app=leetcode.cn id=1518 lang=golang
 *
 * [1518] 换酒问题
 */

// @lc code=start
func numWaterBottles(numBottles int, numExchange int) int {
	result := numBottles
	for numBottles >= numExchange {
		temp := numBottles / numExchange
		result += temp
		numBottles %= numExchange
		numBottles += temp
	}
	return result
}

// @lc code=end
