package leetcode

/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(A []int) bool {
	monotonous := 0 // 0平  1增  2减
	for k := range A {
		if k > 0 {
			if A[k] > A[k-1] {
				if monotonous == 2 {
					return false
				}
				monotonous = 1
			} else if A[k] < A[k-1] {
				if monotonous == 1 {
					return false
				}
				monotonous = 2
			}
		}
	}
	return true
}

// @lc code=end
