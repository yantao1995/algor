package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=937 lang=golang
 *
 * [937] 重新排列日志文件
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	letterStr, numStr := []string{}, []string{}
	for _, v := range logs {
		slice := strings.Split(v, " ")
		if slice[1][0] > 47 && slice[1][0] < 58 { //数字
			numStr = append(numStr, v)
		} else {
			letterStr = append(letterStr, v)
		}
	}
	index := [][]string{}
	for _, v := range letterStr {
		index = append(index, strings.SplitN(v, " ", 2))
	}
	for i := 0; i < len(index)-1; i++ {
		for j := i + 1; j < len(index); j++ {
			if index[i][1] != index[j][1] {
				if index[i][1] > index[j][1] {
					letterStr[i], letterStr[j] = letterStr[j], letterStr[i]
					index[j], index[i] = index[i], index[j]
				}
			} else {
				if index[i][0] > index[j][0] {
					letterStr[i], letterStr[j] = letterStr[j], letterStr[i]
					index[j], index[i] = index[i], index[j]
				}
			}
		}
	}
	return append(letterStr, numStr...)
}

// @lc code=end
