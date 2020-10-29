package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result, tmp := []string{}, []string{}
	digitsSlice := strings.Split(digits, "")
	if len(digitsSlice) == 0 {
		return result
	}
	for _, v1 := range digitsSlice { //分隔值
		temp := m[v1]
		tmp = nil
		for _, v2 := range result { // 返回值
			for _, v3 := range temp { //tmp
				tmp = append(tmp, v2+v3)
			}
		}
		if tmp != nil {
			result = tmp
		} else {
			result = temp
		}
	}
	return result
}

// @lc code=end
