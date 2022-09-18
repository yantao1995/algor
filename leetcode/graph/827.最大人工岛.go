package leetcode

/*
 * @lc app=leetcode.cn id=827 lang=golang
 *
 * [827] 最大人工岛
 */

// @lc code=start
func largestIsland(grid [][]int) int {
	type location struct {
		x, y int
	}
	us := map[location]location{}  //所有连接的1都指向一个location
	countLoc := map[location]int{} //初始location的计数
	zero := []location{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				us[location{i, j}] = location{i, j}
			} else {
				zero = append(zero, location{i, j})
			}
		}
	}
	distinct := map[location]bool{}
	var dfs func(base location, i, j int)
	dfs = func(base location, i, j int) {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid) && grid[i][j] == 1 {
			l := location{i, j}
			if !distinct[l] && us[l] == l {
				distinct[l] = true
				us[l] = base
				countLoc[base]++
				dfs(base, i-1, j)
				dfs(base, i+1, j)
				dfs(base, i, j-1)
				dfs(base, i, j+1)
			}
		}
	}
	tempBase := location{0, 0}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				tempBase = location{i, j}
				if tempBase == us[tempBase] {
					distinct = map[location]bool{}
					dfs(tempBase, i, j)
				}
			}
		}
	}
	maxArea := 1
	board := 0
	getLand := func(i, j int) {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid) && grid[i][j] == 1 {
			tempBase = location{i, j}
			if !distinct[us[tempBase]] {
				distinct[us[tempBase]] = true
				board += countLoc[us[tempBase]]
			}
		}
	}
	for _, zl := range zero {
		board = 0
		distinct = map[location]bool{}
		getLand(zl.x-1, zl.y)
		getLand(zl.x+1, zl.y)
		getLand(zl.x, zl.y-1)
		getLand(zl.x, zl.y+1)
		if maxArea < board+1 {
			maxArea = board + 1
		}
	}
	if len(zero) == 0 {
		maxArea = len(grid) * len(grid)
	}
	return maxArea
}

// @lc code=end

/*
	变种并查集 + dfs
	dfs:和求岛屿面积一样，使用dfs即可拿到岛屿的面积，然后使用 countLoc[location] 存储岛屿初始坐标的面积

	变种并查集,和传统并查集不一样的是此处，将每个岛屿的每块陆地都用并查集指向第一块陆地，便于连接时去重
	比如如下 grid：
			1,0
			1,0
	此时单个岛屿面积为2, 并查集坐标 (0,0) 指向 (0,0), 坐标 (1,0) 指向 (0,0)

	好处是比如遇到如下的 grid：
			1,1,1
			1,0,1
			1,1,1
	用0作为连接时，可以方便的识别出来周围一圈都是同一个岛屿

	然后统计每个0的上下左右的不同岛屿面积之和，找到一个最大值即可。
	如果不存在0，最大岛屿即是矩阵个数之和
*/
