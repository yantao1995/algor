package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1496 lang=golang
 *
 * [1496] 判断路径是否相交
 */

// @lc code=start
func isPathCrossing(path string) bool {
	x, y := 0, 0
	ht := map[string]bool{}
	for k := range path { //N'、'S'、'E' 、 'W'，分别表示向北、向南、向东、向西
		ht[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
		if path[k] == 'N' {
			y++
		}
		if path[k] == 'S' {
			y--
		}
		if path[k] == 'E' {
			x++
		}
		if path[k] == 'W' {
			x--
		}
		if ht[strconv.Itoa(x)+","+strconv.Itoa(y)] {
			return true
		}

	}
	return false
}

// @lc code=end
