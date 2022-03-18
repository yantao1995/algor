package stack

import (
	"algor/achieve/vals"
	"fmt"
)

//Stack ..
type Stack struct {
	StackSlice []vals.AlgorType
	Size       int
	Top        int //栈顶
}

//入
func (s *Stack) Push(str vals.AlgorType) *Stack {
	if s.Size > 0 {
		s.Top++
	}
	if s.Size < len(s.StackSlice) { //复用
		s.StackSlice[s.Top] = str
	} else {
		s.StackSlice = append(s.StackSlice, str)
	}
	s.Size++
	return s
}

//出
func (s *Stack) Pop() vals.AlgorType {
	var rtnVal vals.AlgorType
	if s.Size > 0 {
		//fmt.Printf("%v\t", s.StackSlice[s.Top])
		rtnVal = s.StackSlice[s.Top]
		s.StackSlice[s.Top] = vals.AlgorNilVaule
		if s.Top != 0 {
			s.Top--
		}
		s.Size--
		return rtnVal
	}
	return vals.AlgorNilVaule
}

//获取栈顶元素但不出栈
func (s *Stack) GetTopButNoPop() vals.AlgorType {
	if s.Size > 0 {
		return s.StackSlice[s.Top]
	}
	return vals.AlgorNilVaule
}

//全出
func (s *Stack) PopAll() {
	for s.Top > -1 {
		fmt.Printf("%v\t", s.StackSlice[s.Top])
		s.StackSlice[s.Top] = vals.AlgorNilVaule
		s.Top--
		s.Size--
	}
	s.Top++
	println()
}

//TestStack
func TestStack() {
	stack := new(Stack)
	fmt.Println(stack)
	stack.Push("1").Push("2").Push("3").Push("4")
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	stack.Push("6").Push("7").Push("8").Push("9").Push("10")
	fmt.Println(stack)
	stack.PopAll()
	fmt.Println(stack)
	stack.Push("1").Push("2").Push("3").Push("4")
	fmt.Println(stack)
	fmt.Println(stack.GetTopButNoPop())
	fmt.Println(stack)
	stack.PopAll()
	fmt.Println(stack)
}
