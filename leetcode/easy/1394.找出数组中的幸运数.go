package easy

/*
 * @lc app=leetcode.cn id=1394 lang=golang
 *
 * [1394] 找出数组中的幸运数
 */

// @lc code=start
func findLucky(arr []int) int {
	ht := map[int]int{}
	for k := range arr {
		ht[arr[k]]++
	}
	max := -1
	for k := range ht {
		if k == ht[k] && k > max {
			max = k
		}
	}
	return max
}

// @lc code=end
