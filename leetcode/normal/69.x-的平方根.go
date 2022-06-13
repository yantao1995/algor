package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	mid := x / 2
	lastMid := x
	temp := 0
	for {
		temp = mid - lastMid
		if lastMid > mid {
			temp = lastMid - mid
		}
		if temp >= 2 {
			temp /= 2
		}
		lastMid = mid
		if mid*mid == x || (mid*mid < x && (mid+1)*(mid+1) > x) {
			return mid
		} else if mid*mid > x {
			mid -= temp
		} else {
			mid += temp
		}
	}
}

// @lc code=end
