package leetcode

/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	result := []string{}
	count := 0
	var dp func(i int, stack []byte)
	dp = func(i int, stack []byte) {
		if count == 0 && i == len(s) {
			result = append(result, string(stack))
			return
		}
		for j := i; j < len(s); j++ {
			stack = append(stack, s[j])
			if s[j] == '(' {
				count++
			}
			if s[j] == ')' {
				if count == 0 {
					dp(i+1, stack[:len(stack)-1])
					for k := j - 1; k >= 0; k-- {
						if s[k] == ')' {
							dp(i+1, append(stack[:k], stack[k+1:]...))
						}
					}
				} else {
					count--
				}
			}
		}
	}
	dp(0, []byte{})
	return result
}

// @lc code=end
