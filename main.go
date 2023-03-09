package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "11-{22+(10-5)}"
	// 715 +800 +11250 +43 + 20 = 12828
	fmt.Println(rotationCalc(s)) // 12828.0000000000
}

func rotationCalc(s string) string {
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

type stack struct {
	data []string //用数组表示栈   数组末尾的位置为栈顶
	ptr  int      //栈顶指针 指向栈顶元素
}

func newStack() *stack {
	return &stack{
		data: []string{},
		ptr:  -1,
	}
}

// 入栈
func (s *stack) push(x string) {
	s.ptr++
	if len(s.data) < s.ptr+1 {
		s.data = append(s.data, x)
	}
	s.data[s.ptr] = x
}

// 出栈
func (s *stack) pop() string {
	if s.len() == 0 {
		return ""
	}
	val := s.data[s.ptr]
	s.data = s.data[:s.ptr]
	s.ptr--

	return val
}

// 获取栈顶元素
func (s *stack) empty() bool {
	return s.ptr == -1
}

// 获取栈顶元素
func (s *stack) top() string {
	if s.empty() {
		return ""
	}
	return s.data[s.ptr]
}

// 获取栈的元素个数
func (s *stack) len() int {
	return s.ptr + 1
}
