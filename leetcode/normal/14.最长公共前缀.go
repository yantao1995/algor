package leetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	rtn, temp := "", ""
	for i := 1; ; i++ { //切片左闭右开区间，长度1开始
		if i > len(strs[0]) {
			return rtn
		}
		temp = strs[0][0:i]
		for k := 1; k < len(strs); k++ {
			if i > len(strs[k]) {
				return rtn
			}
			if temp != strs[k][0:i] {
				return rtn
			}
		}
		rtn = temp
	}
}

// @lc code=end
