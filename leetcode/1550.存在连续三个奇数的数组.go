package leetcode

/*
 * @lc app=leetcode.cn id=1550 lang=golang
 *
 * [1550] 存在连续三个奇数的数组
 */

// @lc code=start
func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i := 2; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			if arr[i-1]%2 == 1 && arr[i-2]%2 == 1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
