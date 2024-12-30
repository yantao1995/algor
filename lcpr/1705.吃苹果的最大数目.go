package lcpr

import "slices"

/*
 * @lc app=leetcode.cn id=1705 lang=golang
 * @lcpr version=30204
 *
 * [1705] 吃苹果的最大数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func eatenApples(apples []int, days []int) int {
	result := 0
	stack := []int{}

	for d := 0; ; d++ {

		if d < len(apples) && apples[d] > 0 && days[d] > 0 {
			idx := 0
			for ; idx < len(stack); idx++ {
				if days[d] <= days[stack[idx]]-stack[idx] {
					break
				}
			}
			stack = slices.Insert(stack, idx, d)
		}

		for len(stack) > 0 {
			if d < days[stack[0]]+stack[0] && apples[stack[0]] > 0 {
				result++
				apples[stack[0]]--
				break
			} else {
				stack = stack[1:]
			}
		}

		if len(stack) == 0 && d > len(apples) {
			break
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,5,2]\n[3,2,1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,0,0,0,0,2]\n[3,0,0,0,0,2]\n
// @lcpr case=end

*/

/*

Time Limit Exceeded
62/73 cases passed (N/A)
优化：改成map，减少遍历次数,还是超时
func eatenApples(apples []int, days []int) int {
	result := 0
	mi := map[int]bool{}
	for d := 0; ; d++ {
		minDaysI := 0 //当前最小天数索引
		minDays := -1
		if d < len(apples) {
			mi[d] = true
		}
		for i := range mi {
			if apples[i] > 0 && d-i < days[i] {

				if minDays == -1 ||
					minDays > days[i]+i-d { //剩的天数少就消耗掉

					minDays = days[i] + i - d
					minDaysI = i
				}
			} else {
				delete(mi, i)
			}
		}
		if minDays == -1 {
			if d < len(apples) {
				continue
			}
			break
		}
		apples[minDaysI]--
		result++
	}
	return result
}




思路：每次都从前面列表里取一个满足的并且剩余天数最短的消耗
超时
func eatenApples(apples []int, days []int) int {
	result := 0
	for d := 0; ; d++ {
		minDaysI := 0 //当前最小天数索引
		minDays := -1
		for i := 0; i <= d && i < len(apples); i++ {
			if apples[i] > 0 && d-i < days[i] {

				if minDays == -1 ||
					minDays > days[i]+i-d { //剩的天数少就消耗掉

					minDays = days[i] + i - d
					minDaysI = i
				}
			}
		}
		if minDays == -1 {
			if d < len(apples) {
				continue
			}
			break
		}
		apples[minDaysI]--
		result++
	}
	return result
}
*/
