package leetcode

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	type local struct {
		i, j int
	}
	pacific, atlantic := map[local]bool{}, map[local]bool{}
	for k := range heights {
		pacific[local{k, 0}] = true
		atlantic[local{k, len(heights[0]) - 1}] = true
	}
	for k := range heights[0] {
		pacific[local{0, k}] = true
		atlantic[local{len(heights) - 1, k}] = true
	}
	adj := map[local][]local{}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			adj[local{i, j}] = []local{}
			if i-1 >= 0 && heights[i][j] >= heights[i-1][j] {
				adj[local{i, j}] = append(adj[local{i, j}], local{i - 1, j})
			}
			if i+1 < len(heights) && heights[i][j] >= heights[i+1][j] {
				adj[local{i, j}] = append(adj[local{i, j}], local{i + 1, j})
			}
			if j-1 >= 0 && heights[i][j] >= heights[i][j-1] {
				adj[local{i, j}] = append(adj[local{i, j}], local{i, j - 1})
			}
			if j+1 < len(heights[i]) && heights[i][j] >= heights[i][j+1] {
				adj[local{i, j}] = append(adj[local{i, j}], local{i, j + 1})
			}
		}
	}
	route := map[local]bool{}
	var dfs func(start local, tp bool)
	dfs = func(start local, tp bool) {
		if tp {
			if pacific[start] {
				for k := range route {
					if route[k] {
						pacific[k] = true
					}
				}
				return
			}
		} else {
			if atlantic[start] {
				for k := range route {
					if route[k] {
						atlantic[k] = true
					}
				}
				return
			}
		}

		for k := range adj[start] {
			if !route[adj[start][k]] {
				route[adj[start][k]] = true
				dfs(adj[start][k], tp)
				route[adj[start][k]] = false
			}
		}
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if !pacific[local{i, j}] {
				route = map[local]bool{{i, j}: true}
				dfs(local{i, j}, true)
			}
			if !atlantic[local{i, j}] {
				route = map[local]bool{{i, j}: true}
				dfs(local{i, j}, false)
			}
		}
	}
	result := [][]int{}
	for loc := range pacific {
		if atlantic[loc] {
			result = append(result, []int{loc.i, loc.j})
		}
	}
	return result
}

// @lc code=end

/*
	local 标记坐标
	pacific, atlantic 分别记录可以到达 太平洋 和 大西洋 的左边
	adj 记录当前坐标可达的所有坐标
	dfs 对左标进行深度优先搜索，然后加入到对应可达的太平洋或者大西洋内

	最后遍历 pacific, atlantic ，都存在的就是都可达的

	时间都浪费在了优化上了
	优化1： 第一次 case16超时，然后将一个 dfs拆成两个，分别去找 pacific, atlantic ,
			用tp标记当前找的是什么，否则每次dfs都找两个的话，比较浪费时间
	优化2： 找到之后直接return，不需要继续找到边
	优化3： 如果当前 pacific 没找到就进 pacific 的dfs， atlantic 没找到就进 atlantic 的 dfs
	优化4： 最后一个case超时，一看数据量很大，在遍历route 的时候比较耗时，所以空间换时间，
			每次重进开始dfs的时候就重新生成 route ,即 route = map[local]bool{{i, j}: true}
*/
