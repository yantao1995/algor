package leetcode

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] ||
			(points[i][0] == points[j][0] && points[i][1] <= points[j][1])
	})
	fmt.Println(points)
	xFlag, yFlag := points[1][0] == points[0][0], points[1][1] == points[0][1]
	if xFlag && yFlag {
		return false
	}
	val := float64(points[1][0]-points[0][0]) / float64(points[1][1]-points[0][1])
	for i := 2; i < len(points); i++ {
		if (xFlag && points[i][0] != points[0][0]) ||
			(yFlag && points[i][1] != points[0][1]) {
			return true
		}
		if float64(points[i][0]-points[0][0])/float64(points[i][1]-points[0][1]) != val {
			return true
		}
	}
	return false
}

// @lc code=end

/*
	先对每个数据进行升序排序，然后
	然后比较第一个和第二个是否相等，然后计算斜率
	然后用第三个和第一个算斜率，然后比较斜率是否相等
*/
