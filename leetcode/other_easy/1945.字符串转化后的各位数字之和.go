package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1945 lang=golang
 *
 * [1945] 字符串转化后的各位数字之和
 */

// @lc code=start
func getLucky(s string, k int) int {
	result := ""
	for i := 0; i < len(s); i++ {
		result += strconv.Itoa(int(s[i] - 'a' + 1))
	}
	temp := 0
	for ; len(result) > 1 && k > 0; k-- {
		temp = 0
		for i := 0; i < len(result); i++ {
			temp += int(result[i] - '0')
		}
		result = strconv.Itoa(temp)
	}
	i, _ := strconv.Atoi(result)
	return i
}

// @lc code=end

/*
	直接模拟转化即可
*/
