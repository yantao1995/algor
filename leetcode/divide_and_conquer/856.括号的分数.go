package leetcode

/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 */

// @lc code=start
func scoreOfParentheses(s string) int {
	if len(s) == 2 {
		return 1
	}
	count := 0
	for i := 0; ; i++ {
		if s[i] == '(' {
			count++
		} else {
			count--
			if count == 0 {
				if i == len(s)-1 {
					return 2 * scoreOfParentheses(s[1:len(s)-1])
				}
				return scoreOfParentheses(s[:i+1]) + scoreOfParentheses(s[i+1:])
			}
		}
	}
}

// @lc code=end

/*
	分治：
	左闭右开区间
	（()） 可以为 2* s[1:3]
	 ()（） 可以为 s[0:2] + s[2:4]
*/
