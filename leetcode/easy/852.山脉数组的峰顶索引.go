package easy

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	for k := range arr {
		if k-1 >= 0 && arr[k] < arr[k-1] {
			return k - 1
		}
	}
	return 0
}

// @lc code=end
