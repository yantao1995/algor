package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=1252 lang=golang
 *
 * [1252] 奇数值单元格的数目
 */

// @lc code=start
func oddCells(n int, m int, indices [][]int) int {
	ht := map[string]int{}
	for _, v := range indices {
		for i := 0; i < n; i++ {
			ht[strconv.Itoa(i)+","+strconv.Itoa(v[1])]++
		}
		for j := 0; j < m; j++ {
			ht[strconv.Itoa(v[0])+","+strconv.Itoa(j)]++
		}
	}
	count := 0
	for k := range ht {
		if ht[k]%2 == 1 {
			count++
		}
	}
	return count
}

// @lc code=end
