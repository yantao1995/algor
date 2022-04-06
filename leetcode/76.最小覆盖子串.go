package leetcode

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	m := map[byte]int{}
	for k := range t {
		m[t[k]]++
	}
	minStr := ""
	temp := map[byte]int{}
	flag := true
	for i := len(s) - len(t); i >= 0; i-- {
		temp = map[byte]int{}
		if _, ok := m[s[i]]; ok {
			for j := i; j < len(s); j++ {
				flag = true
				temp[s[j]]++
				for k, v := range m {
					if temp[k] < v {
						flag = false
						break
					}
				}
				if flag && ((j-i+1) < len(minStr) || minStr == "") {
					minStr = s[i : j+1]
					break
				}
			}
		}
	}
	return minStr
}

// @lc code=end
