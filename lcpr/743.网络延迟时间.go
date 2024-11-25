package lcpr

/*
 * @lc app=leetcode.cn id=743 lang=golang
 * @lcpr version=30204
 *
 * [743] 网络延迟时间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	result := -1
	dij := make([][]int, n+1)
	for i := 0; i < len(dij); i++ {
		dij = make([][]int, n+1)
	}
	for _, v := range times {
		dij[v[0]][v[1]] = v[2]
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,1],[2,3,1],[3,4,1]]\n4\n2\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n2\n
// @lcpr case=end

*/
