package leetcode

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	//邻接矩阵  i,j = true
	dp := make([][]bool, numCourses)
	for k := range dp {
		dp[k] = make([]bool, numCourses)
	}
	for k := range prerequisites {
		dp[prerequisites[k][0]][prerequisites[k][1]] = true
	}
	loop := map[int]bool{}
	canLoop := map[int]bool{}
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if loop[i] {
			return false
		}
		loop[i] = true
		if canLoop[i] {
			return true
		}
		canLoop[i] = true
		for j := 0; j < len(dp[i]); j++ {
			if i == j {
				continue
			}
			if dp[i][j] { //连通
				if !dfs(j) {
					return false
				}
				delete(loop, j)
			}
		}
		return true
	}
	for i := 0; i < numCourses; i++ {
		if canLoop[i] {
			continue
		}
		if dp[i][i] || !dfs(i) {
			return false
		}
		delete(loop, i)
	}
	return true
}

// @lc code=end
