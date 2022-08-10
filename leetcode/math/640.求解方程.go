package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=640 lang=golang
 *
 * [640] 求解方程
 */

// @lc code=start
func solveEquation(equation string) string {
	sli := strings.Split(equation, "=")
	left, right := sli[0], sli[1]
	if sli[0][0] != '-' {
		left = "+" + sli[0]
	}
	if sli[1][0] != '-' {
		right = "+" + sli[1]
	}
	temp0, temp1 := 0, 0
	calcTemp := func(str string, vTemp *int) {
		if len(str) == 1 {
			str += "1"
		}
		it, _ := strconv.Atoi(str)
		*vTemp += it
	}
	start := 0
	left2, right2 := 0, 0
	calcStr := func(str string, xTemp, vTemp *int) {
		start = 0
		for i := 0; i < len(str); i++ {
			if i > 0 && (i == len(str)-1 || str[i+1] == '+' || str[i+1] == '-') {
				if str[i] == 'x' {
					calcTemp(str[start:i], xTemp)
				} else {
					calcTemp(str[start:i+1], vTemp)
				}
				start = i + 1
			}
		}
	}
	calcStr(left, &left2, &temp0)
	calcStr(right, &right2, &temp1)

	if right2 > left2 {
		left2, right2, temp0, temp1 = right2, left2, temp1, temp0
	}

	left2 -= right2
	temp1 -= temp0

	if left2 == 0 && temp1 == 0 {
		return "Infinite solutions"
	}
	if left2 == 0 && temp1 != 0 {
		return "No solution"
	}

	return "x=" + strconv.Itoa(temp1/left2)
}

// @lc code=end

/*
	将字符串按 = 号进行分割，
	然后分别合并 x  和  常数项，为了统一计算，首位符号如果不为-，就补充+

	左右分别计算， 再将左右的 x 合并到左边，保证左边的符号为 + ，为了返回时少一次处理

	然后判断分子与分母。得出当前的解是多少个

*/
