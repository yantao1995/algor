package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) int {
	result := 0
	for i := 0; i < len(ops); i++ {
		if i+1 < len(ops) {
			if ops[i+1] == "C" {
				if i+2 < len(ops) {
					ops = append(ops[:i], ops[i+2:]...)
				} else {
					ops = ops[:i]
				}
				i -= 2
			}
		}
	}
	for i := 0; i < len(ops); i++ {
		iInt := 0
		if ops[i] == "+" {
			if i > 1 {
				ops[i] = strPlus682(ops[i-1], ops[i-2])
				iInt, _ = strconv.Atoi(ops[i])
			}
		} else if ops[i] == "D" {
			if i > 0 {
				iInt, _ = strconv.Atoi(ops[i-1])
				iInt *= 2
				ops[i] = strconv.Itoa(iInt)
			}
		} else {
			iInt, _ = strconv.Atoi(ops[i])
		}
		result += iInt
	}
	return result
}

func strPlus682(a, b string) string {
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	return strconv.Itoa(aInt + bInt)
}

// @lc code=end
