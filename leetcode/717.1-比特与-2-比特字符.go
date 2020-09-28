package leetcode

/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			if i > 0 {
				bits = bits[i:]
			}
			break
		}
	}
	if len(bits) == 1 {
		return bits[0] == 0
	}
	if bits[len(bits)-2] == 0 {
		return true
	} else {
		if len(bits) >= 3 {
			if bits[len(bits)-3] == 0 {
				return false
			}
		}

		return true
	}
}

// @lc code=end
