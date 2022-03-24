package leetcode

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	bts := []byte(s)
	var dfs func(si, ei int) string
	dfs = func(si, ei int) string {
		lct, rct := 0, 0
		count := 1
		result := ""
		for i := si; i < ei; i++ {
			if bts[i] == '[' {
				lct++
				if lct > 1 {
					for count > 0 {
						result += dfs(i, ei)
						count--
					}
				}
			} else if bts[i] == ']' {
				rct++
				if rct == lct { //闭环

				}
			} else if bts[i] >= '0' && bts[i] <= '9' {
				if lct == rct && i > si { //闭环
					return string(bts[si:i])
				}
			}
		}
		return result
	}
	return dfs(0, len(bts))
}

// @lc code=end
