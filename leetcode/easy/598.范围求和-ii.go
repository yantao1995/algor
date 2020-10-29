package leetcode

/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	min1 := ops[0][0]
	min2 := ops[0][1]
	for i := 1; i < len(ops); i++ {
		if min1 > ops[i][0] {
			min1 = ops[i][0]
		}
		if min2 > ops[i][1] {
			min2 = ops[i][1]
		}
	}
	return min1 * min2
}

// @lc code=end
