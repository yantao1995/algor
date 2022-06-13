package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
func sortByBits(arr []int) []int {
	ht := map[int]int{}
	for k := range arr {
		b := fmt.Sprintf("%b", arr[k])
		count := 0
		for m := range b {
			if b[m] == '1' {
				count++
			}
		}
		ht[k] = count
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if ht[j] > ht[j+1] || (ht[j] == ht[j+1] && arr[j] > arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				ht[j+1], ht[j] = ht[j], ht[j+1]
			}
		}
	}
	return arr
}

// @lc code=end
