package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	r := 0
	for k := range piles {
		if piles[k] > r {
			r = piles[k]
		}
	}

	canFinish := func(speed int) bool {
		need := 0
		for k := range piles {
			need += int(math.Ceil(float64(piles[k]) / float64(speed)))
		}
		return need <= h && speed > 0
	}

	for l, mid := 0, 0; l < r; {
		mid = (l + r) >> 1
		if canFinish(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

// @lc code=end

/*
	因为速度是随便多少的。而且时间大于piles的长度，所以最优数量应该是0到最大值中的某个值
	所以可以从最大的数量依次递减，一只找到能够完成的数量，但是依次递减，会超时
	所以可以用二分法来找这个数
*/
