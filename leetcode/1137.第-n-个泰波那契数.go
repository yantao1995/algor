package leetcode

/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	ht := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	for i := 3; i <= n; i++ {
		ht[i] = ht[i-1] + ht[i-2] + ht[i-3]
	}
	return ht[n]
}

// @lc code=end
