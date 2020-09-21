package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	sli1 := strings.Split(s, "")
	sli2 := strings.Split(t, "")
	m1 := map[string]int{}
	for k := range sli2 {
		m1[sli2[k]]++
	}
	fmt.Println(m1)
	for k := range sli1 {
		if _, ok := m1[sli1[k]]; ok {
			m1[sli1[k]]--
			if m1[sli1[k]] == 0 {
				delete(m1, sli1[k])
			}
		} else {
			return []byte(sli1[k])[0]
		}
	}
	for k := range m1 {
		return []byte(k)[0]
	}
	return ' '
}

// @lc code=end
