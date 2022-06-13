package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1346 lang=golang
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i]*2 > arr[j] {
				break
			} else if arr[i]*2 == arr[j] || arr[i] == arr[j]*2 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
