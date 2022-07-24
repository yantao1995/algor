package leetcode

/*
 * @lc app=leetcode.cn id=1184 lang=golang
 *
 * [1184] 公交站间的距离
 */

// @lc code=start
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}
	forWard, total := 0, 0
	for i := 0; i < len(distance); i++ {
		if start <= destination {
			if i >= start && i < destination {
				forWard += distance[i]
			}
		} else {
			if i < destination || i >= start {
				forWard += distance[i]
			}
		}
		total += distance[i]
	}
	total -= forWard
	if forWard < total {
		return forWard
	}
	return total
}

// @lc code=end

/*
	forWard 记录正向，区间 [start,destination] 或者 [destination，start] 的距离
	total 记录总距离，总距离减去forWard就是走另一条路的距离
	两者求最小值
*/
