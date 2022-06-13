package leetcode

/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(str string) string {
	bts := []byte(str)
	for k := range bts {
		if bts[k] <= 90 && bts[k] >= 65 {
			bts[k] += 32
		}
	}
	return string(bts)
}

// @lc code=end
