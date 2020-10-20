package leetcode

/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(S string) string {
	bts := []byte(S)
	repateFlag := true
	for repateFlag {
		repateFlag = false
		for i := 1; i < len(bts); i++ {
			if bts[i] == bts[i-1] {
				repateFlag = true
				if i < len(bts)-1 {
					bts = append(bts[:i-1], bts[i+1:]...)
				} else {
					bts = bts[:i-1]
				}
				i -= 2
				if i == -1 {
					i = 0
				}
			}
		}
	}
	return string(bts)
}

// @lc code=end
