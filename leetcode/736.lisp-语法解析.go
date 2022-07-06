package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=736 lang=golang
 *
 * [736] Lisp 语法解析
 */

// @lc code=start
func evaluate(expression string) int {
	slice := []string{}
	start := 0
	for i := 0; i < len(expression); i++ {
		if start < i && expression[start] == ' ' {
			start++
		}
		if expression[i] == ' ' {
			if start != i {
				slice = append(slice, expression[start:i])
			}
			start = i + 1
		} else if expression[i] == '(' {
			slice = append(slice, expression[start:i+1])
			start = i + 1
		} else if expression[i] == ')' {
			if start != i {
				slice = append(slice, expression[start:i])
			}
			slice = append(slice, expression[i:i+1])
			start = i + 1
		}
	}
	i := 0 //slice 全局索引
	var dfs func(m map[string]int) int
	dfs = func(m map[string]int) int {
		result := 0  //当前作用域的结果
		bracket := 0 //当前作用域的未匹配括号数量
		temp := map[string]int{}
		if len(m) > 0 {
			temp = m
		}
		for ; i < len(slice); i++ {
			if slice[i] == "(" {
				bracket++
			} else if slice[i] == ")" {
				bracket--
				if bracket == 0 {
					break
				}
			} else if slice[i] == "let" {
				for i++; slice[i] != "(" && slice[i] != ")"; i += 2 {
					if slice[i+1] != "(" && slice[i+1] != ")" {
						temp[slice[i]], _ = strconv.Atoi(slice[i+1])
					} else {
						if slice[i+1] == "(" {
							i++
							v := dfs(temp)
							i--
							temp[slice[i-1]] = v
							result = v
						} else {
							result = temp[slice[i]]
						}
						break
					}
				}
			} else if slice[i] == "add" {
				v := slice[i+1]
				if v[0] >= '0' && v[0] <= '9' {
					vv, _ := strconv.Atoi(v)
					temp[v] = vv
				}
				if slice[i+1] == "(" {
					i++
					result = dfs(map[string]int{})
					i--
					result += temp[slice[i]]
				} else if slice[i+2] == "(" {
					i += 2
					result = temp[v] + dfs(map[string]int{})
				} else {
					v2 := temp[slice[i+2]]
					v3 := temp[v]
					if _, ok := temp[slice[i+2]]; !ok {
						v2, _ = strconv.Atoi(slice[i+2])
					}
					if _, ok := temp[v]; !ok {
						v3, _ = strconv.Atoi(slice[i+1])
					}
					result = v3 + v2
				}
				i += 2
			} else if slice[i] == "mult" {
				v := slice[i+1]
				if v[0] >= '0' && v[0] <= '9' {
					vv, _ := strconv.Atoi(v)
					temp[v] = vv
				}
				if slice[i+1] == "(" {
					i++
					result = dfs(map[string]int{})
					i--
					result *= temp[slice[i]]
				} else if slice[i+2] == "(" {
					i += 2
					result = temp[v] * dfs(map[string]int{})
				} else {
					v2 := temp[slice[i+2]]
					v3 := temp[v]
					if _, ok := temp[slice[i+2]]; !ok {
						v2, _ = strconv.Atoi(slice[i+2])
					}
					if _, ok := temp[v]; !ok {
						v3, _ = strconv.Atoi(slice[i+1])
					}
					result = v3 * v2
				}
				i += 2
			}
		}
		return result
	}
	return dfs(map[string]int{})
}

// @lc code=end
