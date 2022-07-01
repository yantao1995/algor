package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	slice := []string{}
	last := 0
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' ||
			expression[i] == '-' ||
			expression[i] == '*' {
			slice = append(slice, expression[last:i])
			slice = append(slice, string(expression[i]))
			last = i + 1
		}
	}
	slice = append(slice, expression[last:])
	result := []int{}
	var backtrace func(sli []string)
	backtrace = func(sli []string) {
		if len(sli) == 1 {
			it, _ := strconv.Atoi(sli[0])
			result = append(result, it)
			return
		}
		il, ir, result := 0, 0, 0
		temp := ""
		for i := 1; i < len(sli); i += 2 {
			il, _ = strconv.Atoi(sli[i-1])
			ir, _ = strconv.Atoi(sli[i+1])
			if sli[i] == "+" {
				result = il + ir
			} else if sli[i] == "-" {
				result = il - ir
			} else {
				result = il * ir
			}
			temp = sli[i+1]
			sli[i+1] = strconv.Itoa(result)
			backtrace(append(append([]string(nil), sli[:i-1]...), sli[i+1:]...))
			sli[i+1] = temp
		}
	}
	backtrace(slice)
	return result
}

// @lc code=end
