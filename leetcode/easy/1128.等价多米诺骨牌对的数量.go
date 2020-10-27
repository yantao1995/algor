package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	ht := map[string]int{}
	for _, v1 := range dominoes {
		if v1[1] > v1[0] {
			ht[strconv.Itoa(v1[0])+","+strconv.Itoa(v1[1])]++
		} else {
			ht[strconv.Itoa(v1[1])+","+strconv.Itoa(v1[0])]++
		}
	}
	result := 0
	for k := range ht {
		if ht[k] > 1 {
			ht[k]--
			for ht[k] >= 1 {
				result += ht[k]
				ht[k]--
			}
		}
	}
	return result
}

// @lc code=end
