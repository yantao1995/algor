package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=761 lang=golang
 *
 * [761] 特殊的二进制序列
 */

// @lc code=start
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	strs := []string{}
	count, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			strs = append(strs, "1"+makeLargestSpecial(s[cur+1:i])+"0")
			cur = i + 1
		}
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	return strings.Join(strs, "")
}

// @lc code=end

/*
	分治法
	将1和0分别看成（ 和 ）括号对，问题就变成了，求有效括号并且是最大字典序
	所以只需要将层次更深的括号放前面 ，比如 "(())" 应该在 "()"的前面

	所以在1与0数量相等时，即可构成对应的括号对。采用分治法，将每个匹配个数
		的括号区域进行字典序排序，然后返回即可

*/
