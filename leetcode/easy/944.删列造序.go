package easy

/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */

// @lc code=start
func minDeletionSize(A []string) int {
	everyLength := len(A[0])
	count := 0
	for i := 0; i < everyLength; i++ {
		for k := range A {
			if k > 0 && A[k][i] < A[k-1][i] {
				count++
				break
			}
		}
	}
	return count
}

// @lc code=end
