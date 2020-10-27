package easy

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	bits := []byte(s)
	sort.SliceStable(bits, func(i, j int) bool {
		return bits[i] < bits[j]
	})
	result := []byte{}
	for len(bits) > 0 {
		result = append(result, bits[0])
		bits = bits[1:]
		for i := 0; i < len(bits); i++ {
			if bits[i] != result[len(result)-1] {
				result = append(result, bits[i])
				if i == len(bits)-1 {
					bits = bits[:i]
				} else {
					bits = append(bits[:i], bits[i+1:]...)
				}
				i--
			}
		}
		if len(bits) > 0 {
			result = append(result, bits[len(bits)-1])
			bits = bits[:len(bits)-1]
			for j := len(bits) - 1; j >= 0; j-- {
				if bits[j] != result[len(result)-1] {
					result = append(result, bits[j])
					if j == len(bits)-1 {
						bits = bits[:j]
					} else {
						bits = append(bits[:j], bits[j+1:]...)
					}
				}
			}
		}
	}
	return string(result)
}

// @lc code=end
