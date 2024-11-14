package lcpr

/*
 * @lc app=leetcode.cn id=3249 lang=golang
 * @lcpr version=30204
 *
 * [3249] 统计好节点的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countGoodNodes(edges [][]int) int {
	result := 0
	allM := map[int]bool{}
	m := map[int][]int{}
	mc := map[int]int{}
	mm := map[int][]int{}
	for _, edge := range edges {
		mm[edge[0]] = append(mm[edge[0]], edge[1])
		mm[edge[1]] = append(mm[edge[1]], edge[0])
		allM[edge[0]] = true
		allM[edge[1]] = true
	}
	var line func(head int)
	ready := map[int]bool{}
	line = func(head int) {
		if ready[head] {
			return
		}
		ready[head] = true
		for _, node := range mm[head] {
			if !ready[node] {
				m[head] = append(m[head], node)
				line(node)
			}
		}
	}
	line(0)
	var getCount func(vals []int) int
	getCount = func(vals []int) int {
		rlt := len(vals)
		for _, val := range vals {
			if _, ok := mc[val]; !ok {
				mc[val] = getCount(m[val])
			}
			rlt += mc[val]
		}
		return rlt
	}
	for root := range allM {
		if _, ok := mc[root]; !ok {
			mc[root] += getCount(m[root])
		}
	}
	equals := func(val int, slices ...int) bool {
		for _, slice := range slices {
			if mc[slice] != mc[val] {
				return false
			}
		}
		return true
	}

	for root := range allM {
		if len(m[root]) < 2 ||
			equals(m[root][0], m[root][1:]...) {
			result++
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]\n
// @lcpr case=end

*/

/*
	allM 记录有多少个节点
	m 记录树自上而下的当前节点的紧挨子节点数
	mc 记录当前节点的所有子节点数

	line 自上而下的梳理数据到m
	read 用于禁止双向遍历

	getCount 记录当前子树的所有节点数 （加mc作为缓存快速返回，防止超时）

	equals 判断当前节点的所有子树是否都相同
*/
