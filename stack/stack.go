package stack

import "fmt"

//size - top  = 1
type Stack struct {
	StackSlice []string
	Size       int
	Top        int //栈顶指针
}

//入
func (s *Stack) Push(str string) *Stack {
	if s.Size != 0 {
		if s.Top < s.Size-1 {
			s.Top++
			s.StackSlice[s.Top] = str
			return s
		}
	}
	s.StackSlice = append(s.StackSlice, str)
	s.Size++
	s.Top = len(s.StackSlice) - 1

	return s
}

//出
func (s *Stack) Pop() {
	if s.Top >= 0 {
		fmt.Printf("%s\t", s.StackSlice[s.Top])
		s.Top--
	}
	println()
}

//全出
func (s *Stack) PopAll() {
	for s.Top >= 0 {
		fmt.Printf("%s\t", s.StackSlice[s.Top])
		s.Top--
	}
	println()
}

//TestStack
func TestStack() {
	stack := new(Stack)
	fmt.Println(stack)
	stack.Push("1").Push("2").Push("3").Push("4")
	stack.Pop()
	stack.Push("5")
	stack.Pop()
	stack.PopAll()
}
