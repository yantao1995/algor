package leetcode

/*
 * @lc app=leetcode.cn id=1534 lang=golang
 *
 * [1534] 统计好三元组
 */

// @lc code=start
func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	for k := len(arr) - 1; k > 1; k-- {
		for j := k - 1; j > 0; j-- {
			for i := j - 1; i > -1; i-- {
				if getAbsoult1534(arr[i], arr[j]) <= a &&
					getAbsoult1534(arr[j], arr[k]) <= b &&
					getAbsoult1534(arr[i], arr[k]) <= c {
					result++
				}
			}
		}
	}
	return result
}

func getAbsoult1534(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end
