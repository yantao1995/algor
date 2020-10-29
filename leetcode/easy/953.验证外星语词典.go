package leetcode

/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	ht := map[byte]int{}
	greaterThanPrev := make([]bool, len(words)) //是否比前一个大
	maxLength := 0
	for k := range words {
		if maxLength < len(words[k]) {
			maxLength = len(words[k])
		}
	}
	for k := range order {
		ht[order[k]] = k
	}
	for i := 0; i < maxLength; i++ {
		for k := 1; k < len(words); k++ {
			if !greaterThanPrev[k] {
				if i >= len(words[k-1]) || i >= len(words[k]) {
					if len(words[k-1]) > len(words[k]) {
						return false
					}
				} else {
					if ht[words[k-1][i]] > ht[words[k][i]] {
						return false
					} else {
						if ht[words[k-1][i]] < ht[words[k][i]] {
							greaterThanPrev[k] = true
						}
					}
				}
			}
		}
	}
	return true
}

// @lc code=end
