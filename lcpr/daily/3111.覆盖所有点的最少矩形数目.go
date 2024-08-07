package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=3111 lang=golang
 * @lcpr version=30204
 *
 * [3111] 覆盖所有点的最少矩形数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minRectanglesToCoverPoints(points [][]int, w int) int {
	result := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i, j := 0, 0; i < len(points) && j < len(points); i = j {
		result++
		for j < len(points) && points[j][0]-points[i][0] <= w {
			j++
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1],[1,0],[1,4],[1,8],[3,5],[4,6]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[0,0],[1,1],[2,2],[3,3],[4,4],[5,5],[6,6]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[2,3],[1,2]]\n0\n
// @lcpr case=end

*/

/*
	先排序，因为不统计高度，所以只看x轴长度即可，双指针法，外层循环记录个数，内层循环计算 i ,j 的长度
*/
