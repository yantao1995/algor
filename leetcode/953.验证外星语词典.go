package leetcode

/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	ht := map[byte]int{}
	for k := range order {
		ht[order[k]] = k
	}
	for i := 0; ; i++ {
		for k := 0; k < len(words); k++ {
			if k > 0 {
				if i >= len(words[k-1]) {
					words = append(words[:k-1], words[k:]...)
					k--
					continue
				}
				if ht[words[k-1][i]] > ht[words[k][i]] {
					return false
				} else {
					if ht[words[k-1][i]] == ht[words[k][i]] && len(words[k-1]) > len(words[k]) {
						return false
					}
					words = append(words[:k-1], words[k:]...)
					k--
				}
			}
		}
		if len(words) < 2 {
			return true
		}
	}
}

// @lc code=end
