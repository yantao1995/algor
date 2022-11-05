package leetcode

/*
 * @lc app=leetcode.cn id=1106 lang=golang
 *
 * [1106] 解析布尔表达式
 */

// @lc code=start
func parseBoolExpr(expression string) bool {
	if expression[0] == 't' {
		return true
	} else if expression[0] == 'f' {
		return false
	} else if expression[0] == '!' {
		return !parseBoolExpr(expression[2 : len(expression)-1])
	} else if expression[0] == '&' {
		exprs := expression[2 : len(expression)-1]
		count := 0
		lastIndex := 0
		slice := []string{}
		for i := 0; i < len(exprs); i++ {
			if exprs[i] == '(' {
				count++
			} else if exprs[i] == ')' {
				count--
			} else if exprs[i] == ',' && count == 0 {
				slice = append(slice, exprs[lastIndex:i])
				lastIndex = i + 1
			}
			if i == len(exprs)-1 {
				slice = append(slice, exprs[lastIndex:i+1])
			}
		}
		flag := true
		for i := 0; i < len(slice); i++ {
			flag = flag && parseBoolExpr(slice[i])
		}
		return flag
	} else {
		exprs := expression[2 : len(expression)-1]
		count := 0
		lastIndex := 0
		slice := []string{}
		for i := 0; i < len(exprs); i++ {
			if exprs[i] == '(' {
				count++
			} else if exprs[i] == ')' {
				count--
			} else if exprs[i] == ',' && count == 0 {
				slice = append(slice, exprs[lastIndex:i])
				lastIndex = i + 1
			}
			if i == len(exprs)-1 {
				slice = append(slice, exprs[lastIndex:i+1])
			}
		}
		flag := false
		for i := 0; i < len(slice); i++ {
			flag = flag || parseBoolExpr(slice[i])
		}
		return flag
	}
}

// @lc code=end

/*
	dfs 递归调用，按括号匹配的数量，来判断表达式的作用域
	|()  ,  &()  ,  !() 可以直接分解为：
		 表达式逻辑 + 内部表达式expression[2 : len(expression)-1]
	使用表达式逻辑 ( && 或者 || ) 循环调用 内部表达式即可
*/
