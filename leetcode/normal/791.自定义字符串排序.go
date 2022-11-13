package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=791 lang=golang
 *
 * [791] 自定义字符串排序
 */

// @lc code=start
func customSortString(order string, s string) string {
	m := map[byte]int{}
	for k, v := range order {
		m[byte(v)] = k
	}
	bts := []byte(s)
	sort.Slice(bts, func(i, j int) bool {
		return m[bts[i]] < m[bts[j]]
	})
	return string(bts)
}

// @lc code=end

/*
	获取权重排序即可
*/
