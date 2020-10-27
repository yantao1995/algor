package easy

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	bts := []byte(s)
	for i := 0; i < len(bts); i++ {
		tempJ := i + 1
		if bts[i] != ' ' {
			j := i + 1
			for j < len(bts) && bts[j] != ' ' {
				j++
			}
			j--
			tempJ = j
			for i < j {
				bts[i], bts[j] = bts[j], bts[i]
				i++
				j--
			}
			i = tempJ
		}
	}
	return string(bts)
}

// @lc code=end
