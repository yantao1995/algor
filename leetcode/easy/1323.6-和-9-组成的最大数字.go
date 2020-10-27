package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 */

// @lc code=start
func maximum69Number(num int) int {
	bts := []byte(strconv.Itoa(num))
	for k := range bts {
		if bts[k] == '6' {
			bts[k] = '9'
			break
		}
	}
	temp, _ := strconv.Atoi(string(bts))
	return temp
}

// @lc code=end
