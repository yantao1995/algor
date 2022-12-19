package leetcode

/*
 * @lc app=leetcode.cn id=1971 lang=golang
 *
 * [1971] 寻找图中是否存在路径
 */

// @lc code=start
func validPath(n int, edges [][]int, source int, destination int) bool {
	us := make([]int, n)
	for i := 0; i < n; i++ {
		us[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if us[x] != x {
			us[x] = find(us[x])
		}
		return us[x]
	}
	union := func(a, b int) {
		if a > b {
			a, b = b, a
		}
		us[find(b)] = find(a)
	}
	for _, edge := range edges {
		union(edge[0], edge[1])
	}
	return find(source) == find(destination)
}

// @lc code=end

/*
并查集实现



dfs
Time Limit Exceeded
12/27 cases passed (N/A)

func validPath(n int, edges [][]int, source int, destination int) bool {
	m := map[int][]int{}
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}
	result := false
	var dfs func(s int, route map[int]bool)
	dfs = func(s int, route map[int]bool) {
		if result {
			return
		}
		if s == destination {
			result = true
			return
		}
		for i := 0; i < len(m[s]); i++ {
			if !route[m[s][i]] {
				route[m[s][i]] = true
				dfs(m[s][i], route)
				route[m[s][i]] = false
			}
		}
	}
	dfs(source, map[int]bool{})
	return result
}


map
Time Limit Exceeded
22/27 cases passed (N/A)

func validPath(n int, edges [][]int, source int, destination int) bool {
	m := map[int]bool{source: true}
	for count := 1; count > 0; {
		count = 0
		for _, edge := range edges {
			if m[edge[0]] && !m[edge[1]] {
				m[edge[1]] = true
				count++
			} else if !m[edge[0]] && m[edge[1]] {
				m[edge[0]] = true
				count++
			}
		}
	}
	return m[destination]
}
*/
