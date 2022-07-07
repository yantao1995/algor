package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dictionary []string, sentence string) string {
	m := map[string]bool{}
	for k := range dictionary {
		m[dictionary[k]] = true
	}
	result := make([]string, 0, len(dictionary))
lab:
	for last, i, str := 0, 0, ""; i < len(sentence); i++ {
		if sentence[i] == ' ' || i == len(sentence)-1 {
			if i == len(sentence)-1 {
				str = string(sentence[last:])
			} else {
				str = string(sentence[last:i])
			}
			last = i + 1
			for j := 1; j <= len(str); j++ {
				if m[str[:j]] {
					result = append(result, str[:j])
					continue lab
				}
			}
			result = append(result, str)
		}
	}
	return strings.Join(result, " ")
}

// @lc code=end

/*
	用map存储词根，然后用空格分隔字符串，然后遍历单个字符串找到词根，然后加入结果集
*/
