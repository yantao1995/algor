package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// func longestPalindrome(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	if len(s) == 1 {
// 		return s
// 	}
// 	for lns := len(s); lns > 1; lns-- { //字串长度
// 		for i := 0; i <= len(s)-lns; i++ { //移动游标
// 			flag := true
// 			for j := i; j < lns/2+i; j++ { //字串比较
// 				//fmt.Println("lns", lns, "i", i, "start", j, "end", lns+2*i-j-1)
// 				if s[j] != s[lns+2*i-j-1] { //区间
// 					flag = false
// 					break
// 				}
// 			}
// 			if flag {
// 				return s[i : lns+i]
// 			}
// 		}
// 	}
// 	return string(s[0])
// }

// @lc code=end
