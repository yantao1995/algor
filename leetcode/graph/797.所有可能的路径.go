package leetcode

/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	temp := []int{0}
	var dfs func(target int)
	dfs = func(target int) {
		if target == len(graph)-1 {
			result = append(result, append([]int{}, temp...))
			return
		}
		for k := range graph[target] {
			temp = append(temp, graph[target][k])
			dfs(graph[target][k])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return result
}

// @lc code=end

/*
	因为是有向无环图，所以可以直接dfs找出所有能到达的位置
*/
