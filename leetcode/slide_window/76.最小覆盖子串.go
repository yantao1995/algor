package leetcode

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	m, less := map[byte]int{}, map[byte]bool{}
	for k := range t {
		m[t[k]]++
		less[t[k]] = true
	}
	minStr := ""
	temp := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ { //右边界增加
		if _, ok := m[s[r]]; ok {
			temp[s[r]]++
			if temp[s[r]] >= m[s[r]] {
				if less[s[r]] {
					delete(less, s[r])
				}
				if len(less) == 0 { //移动左边界
					for ; l < r; l++ { //左边界缩小
						if _, ok := m[s[l]]; ok {
							if temp[s[l]] <= m[s[l]] {
								break
							}
							temp[s[l]]--
						}
					}
					if r-l+1 < len(minStr) || minStr == "" {
						minStr = s[l : r+1]
					}
					temp[s[l]]--
					less[s[l]] = true
					l++
				}
			}
		}
	}

	return minStr
}

// @lc code=end

/*
	左右指针 l , r 分别记录当前扫描的前后位置
	m 保存字串的各个字符的数量, less表示当前需要的所有字符
	temp 表示当前扫描区间内的各个字符的数量

	r:移动右指针,扫描的字符数量存里temp，然后达到数量要求时就减少 less 中的对应字符
		如果当前less为空，表示当前区间存在的字符满足
	l:当r满足时，开始移动左区间，缩小到当前区间满足条件的最小值。记录当前的长度，
		移动左指针l,同时记录l弹出的字符到less，并且减少temp的数量。

	直到扫描到r越界
*/
