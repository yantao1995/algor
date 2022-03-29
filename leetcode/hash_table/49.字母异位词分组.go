package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	m := map[string][]string{}
	for k := range strs {
		bts := []byte(strs[k])
		sort.Slice(bts, func(i, j int) bool {
			return bts[i] < bts[j]
		})
		if _, ok := m[string(bts)]; !ok {
			m[string(bts)] = []string{}
		}
		m[string(bts)] = append(m[string(bts)], strs[k])
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// @lc code=end
