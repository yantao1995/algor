package leetcode

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	total, i := 0, 1
	for total+i <= n {
		total += i
		i++
	}
	return i - 1
}

// @lc code=end
