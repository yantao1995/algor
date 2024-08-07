package lcpr

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=682 lang=golang
 * @lcpr version=30204
 *
 * [682] 棒球比赛
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func calPoints(operations []string) int {
	score := make([]int, 0, len(operations))
	total := 0
	for _, v := range operations {
		switch v {
		case "+":
			total += score[len(score)-1] + score[len(score)-2]
			score = append(score, score[len(score)-1]+score[len(score)-2])
		case "D":
			total += score[len(score)-1] * 2
			score = append(score, score[len(score)-1]*2)
		case "C":
			total -= score[len(score)-1]
			score = score[:len(score)-1]
		default:
			vi, _ := strconv.Atoi(v)
			score = append(score, vi)
			total += vi
		}
	}
	return total
}

// @lc code=end

/*
// @lcpr case=start
// ["5","2","C","D","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["5","-2","4","C","D","9","+","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["1"]\n
// @lcpr case=end

*/

/*
按题目要求操作即可
*/
