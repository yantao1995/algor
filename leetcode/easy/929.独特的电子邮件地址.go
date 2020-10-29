package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	ht := map[string]bool{}
	for _, v := range emails {
		slice := strings.Split(v, "@")
		for i := 0; i < len(slice[0]); i++ {
			if slice[0][i] == '+' {
				slice[0] = slice[0][:i]
				break
			}
			if slice[0][i] == '.' {
				if i+1 == len(slice[0]) {
					slice[0] = slice[0][:i]
				} else {
					slice[0] = slice[0][:i] + slice[0][i+1:]
				}
				i--
			}
		}
		ht[slice[0]+"@"+slice[1]] = true
	}
	return len(ht)
}

// @lc code=end
