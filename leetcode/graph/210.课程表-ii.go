package leetcode

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	can := map[int]bool{}
	find := false
	graphMap := map[int][]int{}
	for k := range prerequisites {
		graphMap[prerequisites[k][0]] = append(
			graphMap[prerequisites[k][0]], prerequisites[k][1])
	}
	temp := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if len(graphMap[i]) == 0 {
			can[i] = true
			temp = append(temp, i)
		}
	}
	result := []int{}
	var bfs func(num int, temp []int)
	bfs = func(num int, temp []int) {
		if find {
			return
		}
		for k := range graphMap[num] {
			if !can[graphMap[num][k]] {
				return
			}
		}
		can[num] = true
		if len(temp) == numCourses {
			result = append([]int(nil), temp...)
			find = true
			return
		}
		for i := 0; i < numCourses && !find; i++ {
			if !can[i] {
				bfs(i, append(temp, i))
			}
		}
	}
	for i := 0; i < numCourses && !find; i++ {
		if !can[i] {
			bfs(i, append(temp, i))
		}
	}
	if len(temp) == numCourses {
		return temp
	}
	return result
}

// @lc code=end

/*
	bfs 广度优先遍历
	graphMap 记录每个乘客所依赖的课程
	can 记录当前课程是否可以上课
	find用于找到序列后剪枝，终止退出
	len(temp) == numCourses 用于处理无课程依赖的情况
	此处未采用 找 [a,b] a课程的后续课程 集合[c] 来作为内层循环，若采用，可降低时间复杂度
	如 ： [d,a][e,a] 那么集合c就应该是[d,e]
	因为使用了can[i] 来处理当前，所以最多也就内层进入一次bfs,所以未使用
*/
