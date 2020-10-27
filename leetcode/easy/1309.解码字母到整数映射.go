package easy

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1309 lang=golang
 *
 * [1309] 解码字母到整数映射
 */

// @lc code=start
func freqAlphabets(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && strings.Contains(s[i:i+3], "#") {
			it, _ := strconv.Atoi(s[i : i+2])
			result = append(result, byte(rune(96+it)))
			i += 2
		} else {
			result = append(result, s[i]+48)
		}
	}
	return string(result)
}

// @lc code=end
