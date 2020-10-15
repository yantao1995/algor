package leetcode

/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	Bbytes := []byte(B)
	for i := 0; i < len(B); i++ {
		Bbytes = append(Bbytes[1:], Bbytes[:1]...)
		if string(Bbytes) == A {
			return true
		}
	}
	return false
}

// @lc code=end
