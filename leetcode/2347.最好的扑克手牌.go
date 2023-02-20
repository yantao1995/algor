package leetcode

/*
 * @lc app=leetcode.cn id=2347 lang=golang
 *
 * [2347] 最好的扑克手牌
 */

// @lc code=start
func bestHand(ranks []int, suits []byte) string {
	is1, is2, is3 := true, false, false
	m := map[int]int{}
	for i := 0; i < len(ranks); i++ {
		if is1 && i > 0 && suits[i] != suits[i-1] {
			is1 = false
		}
		m[ranks[i]]++
		if m[ranks[i]] == 3 {
			is2 = true
		}
		if m[ranks[i]] == 2 {
			is3 = true
		}
	}
	if is1 {
		return "Flush"
	} else if is2 {
		return "Three of a Kind"
	} else if is3 {
		return "Pair"
	}
	return "High Card"
}

// @lc code=end

/*
	直接挨个判断状态即可
*/
