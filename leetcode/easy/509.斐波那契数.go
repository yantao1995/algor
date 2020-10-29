package leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	result := make([]int, 31)
	result[1] = 1
	for i := 2; i <= N; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[N]
}

// @lc code=end
