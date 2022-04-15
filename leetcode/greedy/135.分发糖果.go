package leetcode

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	cr := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			cr[i] = 1
		} else {
			cr[i] = 1
			if ratings[i] > ratings[i-1] {
				cr[i] = cr[i-1] + 1
			}
			if ratings[i] < ratings[i-1] {
				for j := i; j-1 >= 0 &&
					ratings[j] < ratings[j-1] &&
					cr[j] >= cr[j-1]; j-- {
					cr[j-1]++
				}
			}
		}
	}
	count := 0
	for k := range cr {
		count += cr[k]
	}
	return count
}

// @lc code=end

/*
	贪心到i个孩子时，给最少的糖果,满足当前最优
	每次都先给1个糖果，然后向左判断。不满足规则就给左边多加一个糖果
*/
