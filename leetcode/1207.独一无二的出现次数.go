package leetcode

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	ht1, ht2 := map[int]int{}, map[int]int{}
	for k := range arr {
		ht1[arr[k]]++
	}
	for k := range ht1 {
		ht2[ht1[k]]++
	}
	for k := range ht2 {
		if ht2[k] > 1 {
			return false
		}
	}
	return true
}

// @lc code=end
