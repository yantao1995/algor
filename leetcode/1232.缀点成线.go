package leetcode

/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	divisorA, divisorB := (coordinates[1][0] - coordinates[0][0]), (coordinates[1][1] - coordinates[0][1])
	for i := 1; i < len(coordinates); i++ {
		if divisorB == 0 {
			if (coordinates[i][1] - coordinates[0][1]) != 0 {
				return false
			}
		} else if divisorA == 0 {
			if (coordinates[i][0] - coordinates[0][0]) != 0 {
				return false
			}
		} else {
			if (float64(divisorA) / float64(divisorB)) != (float64(coordinates[i][0]-coordinates[0][0]) / float64(coordinates[i][1]-coordinates[0][1])) {
				return false
			}
		}
	}
	return true
}

// @lc code=end
