package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=736 lang=golang
 *
 * [736] Lisp 语法解析
 */

// @lc code=start
func evaluate(expression string) int {
	slice := []string{}
	start := 0
	//拆分
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
	//计算
	i := 0 //slice 全局索引
	var dfs func(m map[string]int) int
	dfs = func(m map[string]int) int {
		result := 0  //当前作用域的结果
		bracket := 0 //当前作用域的未匹配括号数量
		temp := map[string]int{}
		if len(m) > 0 {
			for k, v := range m {
				temp[k] = v
			}
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
				i++
				for ; i < len(slice) && slice[i] != "(" && slice[i] != ")"; i++ {
					if slice[i+1] == ")" {
						si, _ := strconv.Atoi(slice[i])
						if _, ok := temp[slice[i]]; ok {
							si = temp[slice[i]]
						}
						result = si
						break
					} else if slice[i+1] == "(" {
						symbol := slice[i]
						i++
						v := dfs(temp)
						temp[symbol] = v
						result = v
					} else {
						if _, ok := temp[slice[i+1]]; ok {
							temp[slice[i]] = temp[slice[i+1]]
						} else {
							temp[slice[i]], _ = strconv.Atoi(slice[i+1])
						}
						i++
					}
				}
				if slice[i] == "(" {
					result = dfs(temp)
				}
			} else {
				tp := slice[i]
				i++
				//前
				v1, v2 := 0, 0
				if slice[i] == "(" {
					v1 = dfs(temp)
				} else {
					v1, _ = strconv.Atoi(slice[i])
					if slice[i][0] >= 'a' && slice[i][0] <= 'z' {
						v1 = temp[slice[i]]
					}
				}
				i++
				//后
				if slice[i] == "(" {
					v2 = dfs(temp)
				} else {
					v2, _ = strconv.Atoi(slice[i])
					if slice[i][0] >= 'a' && slice[i][0] <= 'z' {
						v2 = temp[slice[i]]
					}
				}
				if tp == "add" {
					result = v1 + v2
				} else {
					result = v1 * v2
				}
			}
		}
		return result
	}
	return dfs(map[string]int{})
}

// @lc code=end

/*
	堆屎山  +  dfs
	思路：使用 dfs 来控制变量的作用域
	第一步：先拆分字符串为单个括号，单词
	第二步，使用dfs来控制变量的作用域，遇到 (括号 就进入dfs计算。

	bracket 来判断括号对，如果当前左括号 ( ，全部匹配到了右括号，就跳出当前dfs
	result 记录当前括号对内的运算值
	i  全局变量，记录当前走到的位置
	temp 记录当前括号内的变量值，因为括号内的作用域不一样，所以需要层层向下传递，这样内层修改后，不会影响外层值
	slice 记录当前每一个单词或字符

	分别判断, 因为总是满足表达式的规则，所以默认所有表达式都是合法的

	let ：
		后一个值总是给前一个变量赋值，末尾的表达式或者变量或者值总是为语句的返回值。
		当前语句中，走到slice的i位置
		如果当前i的下一个为 "(" , 那么说明值为一个表达式，直接  当前变量 = dfs
		如果当前i的下一个为 ")" ，那么说明当前为返回值，需要查看是否为变量，是否需要从temp中取值，然后将值交给result

    add 和 mult ：
		逻辑一致，只是运算值的差异是乘还是加,所以用v1,和v2来接收左右值。
		当前语句中，走到slice的i位置
		左右两个值的判断逻辑一样。
		如果当前的i为 "(" ,说明当前是表达式 , 直接进入 dfs
		如果当前的i是值，那么直接取值，如果当前的i是变量，那么从temp中取值

*/
