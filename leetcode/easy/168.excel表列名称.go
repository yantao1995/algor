package leetcode

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(n int) string {
	//10进制转26进制
	mod := []int{}
	for n > 0 {
		mod = append(mod, n%26)
		n /= 26
	}
	//	fmt.Println(mod)
	for i := 0; i < len(mod); i++ {
		flag := false
		if mod[i] == 0 {
			for j := i + 1; j < len(mod); j++ {
				if mod[j] == 1 {
					mod[j] = 26
					flag = true
				} else {
					break
				}
			}
			if flag {
				if i+1 < len(mod) {
					mod = append(mod[:i], mod[i+1:]...)
				} else {
					mod = mod[:i]
				}
			}
		}
	}
	//fmt.Println(mod)
	str := ""
	for i := len(mod) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			if mod[i-1] == 0 {
				if i-1 != 0 {
					i--
					mod[i] = 26
				} else {
					if mod[i] == 1 {
						i--
						mod[i] = 26
					} else {
						mod[i] = 1
						mod[i-1] = 26
					}
				}
			}
		}
		if mod[i] > 0 {
			str += string(rune(mod[i] + 64))
		}
	}
	return str
}

// @lc code=end
