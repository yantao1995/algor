package leetcode

/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return bits[0] == 0
	}
	bits = bits[:len(bits)-1]
	wait := false //前一个是1,在等下一个字符
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			if !wait {
				wait = true
				continue
			}
		} else {
			if !wait {
				continue
			}
		}
		if wait {
			wait = false
		}
	}
	return !wait
}

// @lc code=end
