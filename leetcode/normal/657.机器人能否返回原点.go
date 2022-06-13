package leetcode

/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	ht := map[byte]int{}
	for k := range moves {
		ht[moves[k]]++
	}
	if ht['U'] != ht['D'] {
		return false
	}
	if ht['L'] != ht['R'] {
		return false
	}
	return true
}

// @lc code=end
