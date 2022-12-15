package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1697 lang=golang
 *
 * [1697] 检查边长度限制的路径是否存在
 */

// @lc code=start
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	result := make([]bool, len(queries))
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	fa := make([]int, n)
	for k := range fa {
		fa[k] = k
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	var union func(f, t int)
	union = func(f, t int) {
		fa[find(f)] = find(t)
	}

	for k := range queries {
		queries[k] = append(queries[k], k)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	k := 0
	for _, q := range queries {
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			union(edgeList[k][0], edgeList[k][1])
		}
		result[q[3]] = find(q[0]) == find(q[1])
	}
	return result
}

// @lc code=end

/*
	参考官方题解

	并查集+排序
*/
