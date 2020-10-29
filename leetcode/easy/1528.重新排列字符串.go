package leetcode

/*
 * @lc app=leetcode.cn id=1528 lang=golang
 *
 * [1528] 重新排列字符串
 */

// @lc code=start
func restoreString(s string, indices []int) string {
	bts := make([]byte, len(s))
	for k := range indices {
		bts[indices[k]] = s[k]
	}
	return string(bts)
}

// @lc code=end
