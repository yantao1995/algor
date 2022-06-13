package leetcode

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	sByte := []byte(S)
	for i := 0; i < len(sByte); i++ {
		if sByte[i] == '-' {
			if i == len(sByte)-1 {
				sByte = sByte[:i]
			} else {
				sByte = append(sByte[:i], sByte[i+1:]...)
			}
			i--
		} else {
			if sByte[i] >= 'a' {
				sByte[i] -= 32
			}
		}
	}
	if len(sByte) == 0 {
		return ""
	}
	addLen := len(sByte)/K - 1
	if len(sByte)%K != 0 {
		addLen++
	}
	newByte := make([]byte, addLen+len(sByte))
	newIndex := len(newByte) - 1
	tempK := 0
	for i := len(sByte) - 1; i >= 0; i-- {
		if tempK == K {
			newByte[newIndex] = '-'
			i++
			tempK = 0
		} else {
			newByte[newIndex] = sByte[i]
			tempK++
		}
		newIndex--
	}
	return string(newByte)
}

// @lc code=end
