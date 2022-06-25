package leetcode

/*
	假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
	当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。
	例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
	请计算出粉刷完所有房子最少的花费成本。

*/

func minCost(costs [][]int) int {
	mins := func(a int, b ...int) int {
		for k := range b {
			if b[k] < a {
				a = b[k]
			}
		}
		return a
	}
	for i := 1; i < len(costs); i++ {
		costs[i][0] += mins(costs[i-1][1], costs[i-1][2])
		costs[i][1] += mins(costs[i-1][0], costs[i-1][2])
		costs[i][2] += mins(costs[i-1][0], costs[i-1][1])
	}
	return mins(costs[len(costs)-1][0], costs[len(costs)-1][1:]...)
}

/*
	动态规划
	因为连续的不能使用相同的颜色，所以需要3个dp来记录每个颜色的当前最小值
	所以可以直接使用 costs 当作dp数组，当前颜色就表示当前房子使用当前颜色需要的花费

	可以表示为，当前为 红色 ，可以由前一间的  min(蓝色,绿色) 得到最小花费
*/
