package problems

//栈实现
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
