package leetcode

/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool { //前后各跑一次
	seqFalg := true
	count := 0
	prevIndex, tailIndex := 0, len(s)-1
	for tailIndex > prevIndex {
		if s[prevIndex] != s[tailIndex] {
			if count > 0 {
				seqFalg = false
				break
				//return false
			}
			count++
			if s[tailIndex-1] == s[prevIndex] {
				tailIndex--
			} else if s[prevIndex+1] == s[tailIndex] {
				prevIndex++
			}
			continue
		}
		prevIndex++
		tailIndex--
	}
	if prevIndex == tailIndex && count == 0 {
		seqFalg = true
		//return true
	}
	if tailIndex-prevIndex == 1 && count == 0 {
		seqFalg = false
		//return false
	}
	if seqFalg {
		return true
	}
	count = 0
	prevIndex, tailIndex = 0, len(s)-1
	for tailIndex > prevIndex {
		if s[prevIndex] != s[tailIndex] {
			if count > 0 {
				return false
			}
			count++
			if s[prevIndex+1] == s[tailIndex] {
				prevIndex++
			} else if s[tailIndex-1] == s[prevIndex] {
				tailIndex--
			}
			continue
		}
		prevIndex++
		tailIndex--
	}
	if prevIndex == tailIndex && count == 0 {
		return true
	}
	if tailIndex-prevIndex == 1 && count == 0 {
		return false
	}
	return true
}

// @lc code=end
