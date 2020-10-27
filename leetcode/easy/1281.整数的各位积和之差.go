package easy

/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	bits := []int{}
	for n > 0 {
		bits = append(bits, n%10)
		n /= 10
	}
	product, sum := 1, 0
	for k := range bits {
		product *= bits[k]
		sum += bits[k]
	}
	return product - sum
}

// @lc code=end
