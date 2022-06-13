package leetcode

/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(A []int) bool {
	index := -1
	for k := range A {
		if k > 0 && A[k] == A[k-1] {
			return false
		}
		if k > 0 && A[k] < A[k-1] {
			index = k
			break
		}
	}
	if index-1 == len(A)-1 || index-1 == 0 ||
		index == -1 { //峰顶在第一个或最后一个或没变化
		return false
	}
	for index < len(A) {
		if A[index] == A[index-1] {
			return false
		}
		if A[index] > A[index-1] {
			return false
		}
		index++
	}
	return true
}

// @lc code=end
