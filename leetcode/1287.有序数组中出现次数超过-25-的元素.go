package leetcode

/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	threshold := len(arr) / 4
	lastIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] && i-lastIndex > threshold {
			return arr[i]
		}
		if arr[i] != arr[i-1] {
			lastIndex = i
		}
	}
	if len(arr)-1-lastIndex > threshold {
		return arr[len(arr)-1]
	}
	return arr[lastIndex]
}

// @lc code=end
