package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1807 lang=golang
 *
 * [1807] 替换字符串中的括号内容
 */

// @lc code=start
func evaluate(s string, knowledge [][]string) string {
	m := map[string]string{}
	for _, kl := range knowledge {
		m[kl[0]] = kl[1]
	}
	bts := []byte(s)
	results := []string{}
	lastJ := 0
	for i := 0; i < len(bts); i++ {
		if bts[i] == '(' {
			for j := i + 1; j < len(bts); j++ {
				if bts[j] == ')' {
					results = append(results, string(bts[lastJ:i]))
					if v, ok := m[string(bts[i+1:j])]; ok {
						results = append(results, v)
					} else {
						results = append(results, "?")
					}
					lastJ = j + 1
					i = j
					break
				}
			}
		}
	}
	if lastJ != len(bts) {
		results = append(results, string(bts[lastJ:]))
	}
	return strings.Join(results, "")
}

// @lc code=end

/*
再优化，将每个片段都放到results里面，然后最后在拼接



双指针i和j分别指向每一个'(' 和 ')',然后向后移动，依次判断和替换
内存消耗的比较多
func evaluate(s string, knowledge [][]string) string {
	m := map[string]string{}
	for _, kl := range knowledge {
		m[kl[0]] = kl[1]
	}
	bts := []byte(s)
	temp := ""
	for i := 0; i < len(bts); i++ {
		if bts[i] == '(' {
			for j := i + 1; j < len(bts); j++ {
				if bts[j] == ')' {
					if v, ok := m[string(bts[i+1:j])]; ok {
						temp = string(bts[j+1:])
						bts = append(append(bts[:i], v...), temp...)
						i += len(v) - 1
					} else {
						bts = append(append(bts[:i], '?'), bts[j+1:]...)
					}
					break
				}
			}
		}
	}
	return string(bts)
}


Time Limit Exceeded
98/105 cases passed (N/A)
用库函数strings做字符串替换，然后拼接

func evaluate(s string, knowledge [][]string) string {
	for _, kl := range knowledge {
		s = strings.ReplaceAll(s, "("+kl[0]+")", kl[1])
	}
	idx := strings.IndexByte(s, '(')
	for idx >= 0 {
		s = s[:idx] + "?" + s[strings.IndexByte(s, ')')+1:]
		idx = strings.IndexByte(s, '(')
	}
	return s
}
*/
