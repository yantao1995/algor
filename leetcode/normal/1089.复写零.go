package leetcode

/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			i++
			arr[i] = 0
		}
	}
}

// @lc code=end
