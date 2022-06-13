package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func dayOfYear(date string) int {
	dateSlice := strings.Split(date, "-")
	dateSliceInt := []int{0, 0, 0}
	for k := range dateSlice {
		dateSliceInt[k], _ = strconv.Atoi(dateSlice[k])
	}
	ht := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	if (dateSliceInt[0]%4 == 0 && dateSliceInt[0]%100 != 0) || dateSliceInt[0]%400 == 0 {
		ht[2] = 29
	}
	sequence := 0
	for i := 1; i <= dateSliceInt[1]; i++ {
		if i < dateSliceInt[1] {
			sequence += ht[i]
		} else {
			sequence += dateSliceInt[2]
		}
	}
	return sequence
}

// @lc code=end
