package leetcode

/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 */

// @lc code=start
func possibleBipartition(n int, dislikes [][]int) bool {
	m := map[int][]int{}
	for _, dis := range dislikes {
		m[dis[0]] = append(m[dis[0]], dis[1])
		m[dis[1]] = append(m[dis[1]], dis[0])
	}
	sz := make([]int, n+1)
	signal := true
	var dfs func(key, val int)
	dfs = func(key, val int) {
		if !signal {
			return
		}
		sz[key] = val
		for _, dis := range m[key] {
			if sz[dis] == val {
				signal = false
				return
			}
			if sz[dis] == 0 {
				dfs(dis, 0-val)
			}
		}
	}
	for i, val := range sz {
		if !signal {
			break
		}
		if val == 0 {
			dfs(i, 1)
		}
	}
	return signal
}

// @lc code=end

/*
	dfs 对每一层进行染色，sz存储所有需要染色的对象(索引0除外)
		dfs只对未染色对象(值为0)进行染色，若当前层是1，那么下一层就是-1
*/
