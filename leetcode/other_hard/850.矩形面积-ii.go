package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=850 lang=golang
 *
 * [850] 矩形面积 II
 */

// @lc code=start
func rectangleArea(rectangles [][]int) int {
	lines := make([]int, 0, len(rectangles)*2)
	for k := range rectangles {
		lines = append(lines, rectangles[k][0], rectangles[k][2])
	}
	sort.Ints(lines)
	var mod int = 1e9 + 7
	result := 0
	innerLines := [][2]int{}
	low, high := -1, -1
	height := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] == lines[i-1] {
			continue
		}
		innerLines = [][2]int{}
		for j := 0; j < len(rectangles); j++ {
			if rectangles[j][0] <= lines[i-1] && rectangles[j][2] >= lines[i] {
				innerLines = append(innerLines, [2]int{rectangles[j][1], rectangles[j][3]})
			}
		}
		sort.Slice(innerLines, func(i, j int) bool {
			return innerLines[i][0] < innerLines[j][0] ||
				(innerLines[i][0] == innerLines[j][0] &&
					innerLines[i][1] <= innerLines[j][1])
		})
		height = 0
		low, high = -1, -1
		for j := 0; j < len(innerLines); j++ {
			height += high - low
			if innerLines[j][0] >= high {
				low, high = innerLines[j][0], innerLines[j][1]
			} else if innerLines[j][1] > high {
				low = high
				high = innerLines[j][1]
			} else {
				low = high
			}
		}
		height += high - low
		result += (height) * (lines[i] - lines[i-1])
		result %= mod
	}
	return result
}

// @lc code=end

/*
	扫描线
	先按X轴，对每一个坐标点进行竖向切割，分隔成一个一个的区间
	然后对每一个区间，进行 y 轴切割，记录 y轴的有效长度，然后与x轴宽度计算出区间内的切割面积。
	最后将每一个区间累加即可。

*/
