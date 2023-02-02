package leetcode

/*
 * @lc app=leetcode.cn id=1129 lang=golang
 *
 * [1129] 颜色交替的最短路径
 */

// @lc code=start
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	r, b := make([][]int, n), make([][]int, n)
	for _, edge := range redEdges {
		r[edge[0]] = append(r[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		b[edge[0]] = append(b[edge[0]], edge[1])
	}
	dp := [3][]int{} //0red   1blue   2result
	for k := range dp {
		dp[k] = make([]int, n)
	}
	route0, route1 := make([]bool, n), make([]bool, n)
	route0[0], route1[0] = true, true
	var dfs func(flag bool, i int)
	dfs = func(flag bool, i int) {
		if flag {
			for _, next := range r[i] {
				if !route1[next] {
					if dp[1][next] == 0 || dp[1][next] > dp[0][i]+1 {
						dp[1][next] = dp[0][i] + 1
						route1[next] = true
						dfs(!flag, next)
						route1[next] = false
					}
				}
			}
		} else {
			for _, next := range b[i] {
				if !route0[next] {
					if dp[0][next] == 0 || dp[0][next] > dp[1][i]+1 {
						dp[0][next] = dp[1][i] + 1
						route0[next] = true
						dfs(!flag, next)
						route0[next] = false
					}

				}
			}
		}
	}
	dfs(true, 0)
	dfs(false, 0)
	for i := 1; i < n; i++ {
		dp[2][i] = dp[0][i]
		if dp[0][i] == 0 || (dp[1][i] > 0 && dp[1][i] < dp[0][i]) {
			dp[2][i] = dp[1][i]
		}
		if dp[2][i] == 0 {
			dp[2][i] = -1
		}
	}
	return dp[2]
}

// @lc code=end

/*
dfs
dp[0]和dp[1] 分别记录 red 和 blue 到达当前点的最短路径
dp[2] 记录结果

优化  使用数组来存储路径访问结果

剪枝  当前节点的下一个结点 如果找到更短路径，才继续向下
if dp[0][next] == 0 || dp[0][next] > dp[1][i]+1 {}


Accepted
90/90 cases passed (8 ms)
Your runtime beats 90.91 % of golang submissions
Your memory usage beats 90.91 % of golang submissions (5.1 MB)


优化前  双 27%


内存优化前，使用的map
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	r, b := map[int][]int{}, map[int][]int{}
	for _, edge := range redEdges {
		r[edge[0]] = append(r[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		b[edge[0]] = append(b[edge[0]], edge[1])
	}
	dp := [3][]int{} //0red   1blue   2result
	for k := range dp {
		dp[k] = make([]int, n)
	}

	route0, route1 := map[int]bool{0: true}, map[int]bool{0: true}
	var dfs func(flag bool, i int)
	dfs = func(flag bool, i int) {
		if flag {
			for _, next := range r[i] {
				if !route1[next] {
					if dp[1][next] == 0 || dp[1][next] > dp[0][i]+1 {
						dp[1][next] = dp[0][i] + 1
						route1[next] = true
						dfs(!flag, next)
						route1[next] = false
					}
				}
			}
		} else {
			for _, next := range b[i] {
				if !route0[next] {
					if dp[0][next] == 0 || dp[0][next] > dp[1][i]+1 {
						dp[0][next] = dp[1][i] + 1
						route0[next] = true
						dfs(!flag, next)
						route0[next] = false
					}

				}
			}
		}
	}
	dfs(true, 0)
	dfs(false, 0)

	for i := 1; i < n; i++ {
		dp[2][i] = dp[0][i]
		if dp[0][i] == 0 || (dp[1][i] > 0 && dp[1][i] < dp[0][i]) {
			dp[2][i] = dp[1][i]
		}
		if dp[2][i] == 0 {
			dp[2][i] = -1
		}
	}
	return dp[2]
}

*/
