package leetcode

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	if word[0] > 90 {
		for i := 1; i < len(word); i++ {
			if word[i] <= 90 {
				return false
			}
		}
		return true
	}
	if len(word) > 1 {
		if word[1] > 90 {
			for i := 1; i < len(word); i++ {
				if word[i] <= 90 {
					return false
				}
			}
			return true
		} else {
			for i := 1; i < len(word); i++ {
				if word[i] > 90 {
					return false
				}
			}
			return true
		}
	}
	return true
}

// @lc code=end
