package easy

import "strings"

/*
 * @lc app=leetcode.cn id=1592 lang=golang
 *
 * [1592] 重新排列单词间的空格
 */

// @lc code=start
func reorderSpaces(text string) string {
	slice := strings.Split(text, " ")
	spaceCount := 0
	for k := range text {
		if text[k] == ' ' {
			spaceCount++
		}
	}
	temp := []string{}
	for k := range slice {
		if slice[k] != "" {
			temp = append(temp, slice[k])
		}
	}
	result := ""
	if len(slice)-len(temp) != 0 && len(temp) == 1 {
		result = temp[0]
		difLen := len(text) - len(result)
		for difLen > 0 {
			result += " "
			difLen--
		}
		return result
	}
	if len(temp) == 1 && len(slice) == len(temp) {
		return text
	}

	count := spaceCount / ((len(temp)) - 1)

	for k := range temp {
		result += temp[k]
		t := count
		for t > 0 && k != len(temp)-1 {
			result += " "
			t--
		}
		if k == len(temp)-1 {
			difLen := len(text) - len(result)
			for difLen > 0 {
				result += " "
				difLen--
			}
		}
	}
	return result
}

// @lc code=end
