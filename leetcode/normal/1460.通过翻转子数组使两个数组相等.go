package leetcode

/*
 * @lc app=leetcode.cn id=1460 lang=golang
 *
 * [1460] 通过翻转子数组使两个数组相等
 */

// @lc code=start
func canBeEqual(target []int, arr []int) bool {
	ht1, ht2 := map[int]int{}, map[int]int{}
	for k := range target {
		ht1[target[k]]++
	}
	for k := range arr {
		ht2[arr[k]]++
	}
	for k := range ht1 {
		if ht1[k] != ht2[k] {
			return false
		}
	}
	return true
}

// @lc code=end
