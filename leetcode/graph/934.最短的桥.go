package leetcode

/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */

// @lc code=start
func shortestBridge(grid [][]int) int {
	g1, g2 := [][2]int{}, [][2]int{}
	distinct := map[[2]int]bool{}
	var dfs func(i, j int, isg1 bool)
	dfs = func(i, j int, isg1 bool) {
		if i < 0 || i >= len(grid) ||
			j < 0 || j >= len(grid) ||
			distinct[[2]int{i, j}] ||
			grid[i][j] == 0 {
			return
		}
		temp := [2]int{i, j}
		if isg1 {
			g1 = append(g1, temp)
		} else {
			g2 = append(g2, temp)
		}
		distinct[temp] = true
		dfs(i-1, j, isg1)
		dfs(i+1, j, isg1)
		dfs(i, j-1, isg1)
		dfs(i, j+1, isg1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 1 && !distinct[[2]int{i, j}] {
				dfs(i, j, len(distinct) == 0)
			}
		}
	}
	getAbs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	getLength := func(a, b [2]int) int {
		return getAbs(a[0], b[0]) + getAbs(a[1], b[1])
	}

	minLength := len(grid) * len(grid)
	for k := range g1 {
		for k2 := range g2 {
			minLength = min(minLength, getLength(g1[k], g2[k2]))
		}
	}
	return minLength - 1
}

// @lc code=end

/*
	先用 dfs 找到两座岛的所有坐标，
	然后对两座岛进行坐标距离连通判断，
	找到最小的一个距离即可（1到5 的距离是4，但是连通只需要接通 2，3，4,也就是3，所以需要 minLength-1）
*/
