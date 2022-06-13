package leetcode

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
// 不能用 要求不使用乘法、除法和 mod 运算符。
// @lc code=start
func divide(dividend int, divisor int) int {

	signA, signB := 1, 1
	result := 0
	if dividend < 0 {
		dividend = 0 - dividend
		signA = -1
	}
	if divisor < 0 {
		divisor = 0 - divisor
		signB = -1
	}
	if divisor > divisor {
		return 0
	}
	tempDivisor := divisor
	tempResult := 1
	for dividend >= divisor { // 被除数
		if dividend-tempDivisor < tempDivisor {
			tempDivisor = divisor
			tempResult = 1
		} else {
			tempDivisor += tempDivisor
			tempResult += tempResult
		}
		result += tempResult
		dividend -= tempDivisor
	}
	if signA+signB == 0 { //异号 结果为负数
		return 0 - result
	}
	if result > 1<<31-1 || result < -1<<31 {
		return 1<<31 - 1
	}
	return result
}

// @lc code=end
