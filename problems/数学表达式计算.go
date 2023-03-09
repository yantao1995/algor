package problems

import "strconv"

//数学表达式计算 需要是正确的数学表达式
func mathExpressionCalc(s string) string {
	stack := newStack()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{', '+', '-', '*', '/': //可以直接入栈的符号
			stack.push(string(s[i])) //转换成 string  入栈
		case ')', ']', '}': //反括号 需要运算得到括号内的结果
			for {
				val := stack.pop()     //括号前一个值
				operate := stack.pop() //运算符
				if operate == "(" || operate == "[" || operate == "{" {
					stack.push(val)
					break
				}
				last := stack.pop() //上一个值
				current := calc(last, val, operate)
				stack.push(current) //当前值入栈
			}

		default: //数字
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			numStr := s[i:j] // 这个字符串是数字
			i = j - 1        //移动到数字的下一个索引
			stack.push(numStr)
		}
		//处理当前栈
		for stack.len() > 2 { //a值  b值 运算符  数量大于3 可以判断操作处理  保证最后只剩一个元素
			stackLen := stack.len()
			current := stack.pop()
			if current != "+" && current != "-" && current != "*" && current != "/" &&
				current != "(" && current != "[" && current != "{" { //当前不是运算符或特殊符号
				operate := stack.pop()
				if operate == "*" || operate == "/" || //第二个是运算符才计算
					((operate == "+" || operate == "-") && i+1 < len(s) && (s[i+1] == '+' || s[i+1] == '-')) ||
					i+1 >= len(s) {
					last := stack.pop()
					current = calc(last, current, operate)
				} else {
					stack.push(operate) //归还运算符
				}
			} else { //当前是运算符或特殊符号
				stack.push(current)
				break
			}
			stack.push(current)
			if stackLen == stack.len() { //当前栈无需处理就退出
				break
			}
		}
	}
	return stack.pop()
}

// 参数分别为 第一个数，第二个数，运算符
func calc(a, b, operate string) string {
	aFloat, _ := strconv.ParseFloat(a, 64) //转换成 float
	bFloat, _ := strconv.ParseFloat(b, 64)
	result := 0.0
	switch operate {
	case "+":
		result = aFloat + bFloat
	case "-":
		result = aFloat - bFloat
	case "*":
		result = aFloat * bFloat
	case "/":
		result = aFloat / bFloat
	}
	return strconv.FormatFloat(result, 'f', 10, 64) //转换成字符串
}
