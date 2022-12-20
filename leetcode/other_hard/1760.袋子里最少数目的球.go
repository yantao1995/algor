package leetcode

/*
 * @lc app=leetcode.cn id=1760 lang=golang
 *
 * [1760] 袋子里最少数目的球
 */

// @lc code=start
func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 0
	for _, num := range nums {
		if right < num {
			right = num
		}
	}
	result := 0
	for left <= right {
		mid := (left + right) / 2
		opts := 0
		for _, num := range nums {
			opts += (num - 1) / mid
		}
		if opts <= maxOperations {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

// @lc code=end

/*
	二分查找，将一个数分解出的最小数就是  num/2
	暴力找出这个数即可
*/
