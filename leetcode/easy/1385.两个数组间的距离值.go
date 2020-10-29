package leetcode

/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0
	for k := range arr1 {
		for m := range arr2 {
			dist := arr1[k] - arr2[m]
			if dist < 0 {
				dist = 0 - dist
			}
			if dist <= d {
				goto come
			}
		}
		count++
	come:
	}
	return count
}

// @lc code=end
