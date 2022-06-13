package leetcode

/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(A []int, K int) int {
	min, max := 2<<32, -1<<32-1
	for k := range A {
		if A[k] > max {
			max = A[k]
		}
		if A[k] < min {
			min = A[k]
		}
	}
	if 2*K > max-min {
		return 0
	}
	return max - min - 2*K
}

// @lc code=end
