package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	bts := []byte(strconv.Itoa(num))
	var bti int
	for i := 0; i < len(bts); i++ {
		bti = i
		for j := len(bts) - 1; j > i; j-- {
			if bts[j] > bts[bti] {
				bti = j
			}
		}
		if bts[i] < bts[bti] {
			bts[i], bts[bti] = bts[bti], bts[i]
			bti, _ = strconv.Atoi(string(bts))
			return bti
		}
	}
	return num
}

// @lc code=end

/*
	遍历序列，如果位置比当前数靠右，有比当前数字大的数，就找一个最大的来替换即可
*/
