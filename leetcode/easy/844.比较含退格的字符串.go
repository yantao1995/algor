package easy

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	Sbyte, Tbyte := []byte(S), []byte(T)
	for k := 0; k < len(Sbyte); k++ {
		if Sbyte[k] == '#' {
			if k+1 < len(Sbyte) {
				if k-1 >= 0 {
					Sbyte = append(Sbyte[:k-1], Sbyte[k+1:]...)
					k -= 2
				} else {
					Sbyte = Sbyte[k+1:]
					k--
				}
			} else {
				if k-1 >= 0 {
					Sbyte = Sbyte[:k-1]
					k -= 2
				} else {
					Sbyte = Sbyte[:k]
					k--
				}
			}

		}
	}
	for k := 0; k < len(Tbyte); k++ {
		if Tbyte[k] == '#' {
			if k+1 < len(Tbyte) {
				if k-1 >= 0 {
					Tbyte = append(Tbyte[:k-1], Tbyte[k+1:]...)
					k -= 2
				} else {
					Tbyte = Tbyte[k+1:]
					k--
				}
			} else {
				if k-1 >= 0 {
					Tbyte = Tbyte[:k-1]
					k -= 2
				} else {
					Tbyte = Tbyte[:k]
					k--
				}
			}
		}
	}
	return string(Sbyte) == string(Tbyte)
}

// @lc code=end
