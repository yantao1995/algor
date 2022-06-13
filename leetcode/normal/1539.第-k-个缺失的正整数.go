package leetcode

/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	count, index := 0, 0
	for i := 1; ; i++ {
		if index < len(arr) {
			if arr[index] != i {
				count++
			} else {
				index++
			}
		} else {
			count++
		}
		if count == k {
			return i
		}
	}
}

// @lc code=end
