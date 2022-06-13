package leetcode

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 颠倒字符串中的单词
 */

// @lc code=start
// func reverseWords(s string) string {
// 	slice := strings.Split(s, " ")
// 	result := ""
// 	for i := len(slice) - 1; i >= 0; i-- {
// 		if slice[i] != "" {
// 			result += slice[i]
// 			if i > 0 {
// 				result += " "
// 			}
// 		}
// 	}
// 	if result[len(result)-1] == ' ' {
// 		return result[:len(result)-1]
// 	}
// 	return result
// }

// @lc code=end

/*
	按空格分隔，然后累加，注意末尾的空格
*/
