package stack

import "fmt"

const popNilVaule = "nil"

type stackType string

//Stack ..
type Stack struct {
	StackSlice []stackType
	Size       int
	Top        int //栈顶
}

//入
func (s *Stack) Push(str stackType) *Stack {
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
func (s *Stack) Pop() *Stack {
	if s.Top >= 0 {
		fmt.Printf("%v\t", s.StackSlice[s.Top])
		s.StackSlice[s.Top] = popNilVaule
		if s.Size != 0 {
			s.Top--
		}
		s.Size--
		println()
	} else {
		println(popNilVaule)
	}
	return s
}

//全出
func (s *Stack) PopAll() {
	for s.Top > -1 {
		fmt.Printf("%v\t", s.StackSlice[s.Top])
		s.StackSlice[s.Top] = popNilVaule
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
	stack.Pop().Pop().Pop()
	fmt.Println(stack)
	stack.Push("6").Push("7").Push("8").Push("9").Push("10")
	fmt.Println(stack)
	stack.PopAll()
	fmt.Println(stack)
	stack.Push("1").Push("2").Push("3").Push("4")
	fmt.Println(stack)
	stack.PopAll()
	fmt.Println(stack)
}
