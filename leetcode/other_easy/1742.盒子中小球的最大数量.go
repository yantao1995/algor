package leetcode

/*
 * @lc app=leetcode.cn id=1742 lang=golang
 *
 * [1742] 盒子中小球的最大数量
 */

// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
	result := 0
	getBitSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return sum
	}
	box := map[int]int{}
	sum := 0
	for lowLimit <= highLimit {
		sum = getBitSum(lowLimit)
		box[sum]++
		if box[sum] > result {
			result = box[sum]
		}
		lowLimit++
	}
	return result
}

// @lc code=end

/*
	用map存每个位数之和的相同的个数，然后找到最大值即可
*/
