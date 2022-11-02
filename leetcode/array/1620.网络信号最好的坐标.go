package leetcode

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1620 lang=golang
 *
 * [1620] 网络信号最好的坐标
 */

// @lc code=start
func bestCoordinate(towers [][]int, radius int) []int {
	if len(towers) == 1 && towers[0][2] == 0 {
		return []int{0, 0}
	}
	locationMap := map[[2]int]int{} //记录左边的信号叠加
	getSign := func(q, x1, y1, x2, y2 int) int {
		distance := math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
		if distance > float64(radius) {
			return 0
		}
		return int(float64(q) /
			(1 + distance))
	}
	for _, tower := range towers {
		i := 0
		if tower[0]-radius > i {
			i = tower[0] - radius
		}
		for ; i <= tower[0]+radius; i++ {
			j := 0
			if tower[1]-radius > j {
				j = tower[1] - radius
			}
			for ; j <= tower[1]+radius; j++ {
				if tower[0] == i && tower[1] == j {
					locationMap[[2]int{i, j}] += tower[2]
				} else {
					locationMap[[2]int{i, j}] += getSign(tower[2], tower[0], tower[1], i, j)
				}
			}
		}
	}
	keys := make([][2]int, 0, len(locationMap))
	for k := range locationMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return locationMap[keys[i]] > locationMap[keys[j]] ||
			(locationMap[keys[i]] == locationMap[keys[j]] && ((keys[i][0] < keys[j][0]) ||
				(keys[i][0] == keys[j][0] && keys[i][1] < keys[j][1])))

	})
	return keys[0][:2]
}

// @lc code=end

/*
	直接依次遍历信号塔，然后找到信号塔覆盖范围的所有整数点，并对整数点的信号值进行累加

	需要注意的是：信号塔的覆盖范围是圆，而不是矩形。所以在得到点之间的距离的时候，需要判断是不是在圆内。
*/
