package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=593 lang=golang
 *
 * [593] 有效的正方形
 */

// @lc code=start
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	verify := func(slice ...int) bool {
		sort.Ints(slice)
		return slice[0] == slice[1] &&
			slice[0]+slice[1] == slice[2] &&
			slice[0] != slice[2]
	}
	handlePoint := func(t1, t2, t3 []int) bool {
		return verify((t1[0]-t2[0])*(t1[0]-t2[0])+(t1[1]-t2[1])*(t1[1]-t2[1]),
			(t1[0]-t3[0])*(t1[0]-t3[0])+(t1[1]-t3[1])*(t1[1]-t3[1]),
			(t3[0]-t2[0])*(t3[0]-t2[0])+(t3[1]-t2[1])*(t3[1]-t2[1]))
	}
	return handlePoint(p1, p2, p3) &&
		handlePoint(p1, p2, p4) &&
		handlePoint(p1, p3, p4) &&
		handlePoint(p2, p3, p4)
}

// @lc code=end

/*
	思路1：正方形内，任意3个顶点必然组成一个 等腰直角三角形

	思路2：找到每条边的对应关系，然后求相邻边的斜率相乘是否为-1，但是找每个点的关系篇幅比较长

*/
