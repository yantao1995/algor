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
	var dac func(sli []string) []int
	dac = func(sli []string) []int {
		result := []int{}
		if len(sli) == 1 {
			it, _ := strconv.Atoi(sli[0])
			return []int{it}
		}
		for i := 1; i < len(sli); i += 2 {
			d1, d2 := dac(append([]string(nil), sli[:i]...)), dac(append([]string(nil), sli[i+1:]...))
			for _, i1 := range d1 {
				for _, i2 := range d2 {
					if sli[i] == "+" {
						result = append(result, i1+i2)
					} else if sli[i] == "-" {
						result = append(result, i1-i2)
					} else {
						result = append(result, i1*i2)
					}
				}
			}
		}
		return result
	}
	return dac(slice)
}

// @lc code=end

/*
	分治法
	先将字符串拆成数字和运算符
	然后按运算符进行分治拆分，拆成单个的结果集
	然后向上递归的时候，对结果集按运算符进行组合运算
*/
