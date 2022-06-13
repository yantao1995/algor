package leetcode

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	ht := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := ht[s[i]]; !ok {
			for k := range ht {
				if ht[k] == t[i] {
					return false
				}
			}
			ht[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}

// @lc code=end
