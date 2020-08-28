package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
// func isValid(s string) bool {
// 	if len(s)%2 != 0 {
// 		return false
// 	}
// 	stack := Stack{
// 		StackSlice: []byte{},
// 		Size:       0,
// 		Top:        0,
// 	}
// 	for k := range s {
// 		if stack.Size > 0 {
// 			if s[k]-stack.GetTopButNoPop() == 1 || s[k]-stack.GetTopButNoPop() == 2 {
// 				stack.Pop()
// 			} else {
// 				stack.Push(s[k])
// 			}
// 		} else {
// 			stack.Push(s[k])
// 		}
// 	}
// 	if stack.Size > 0 {
// 		return false
// 	}
// 	return true
// }

// //Stack ..
// type Stack struct {
// 	StackSlice []byte
// 	Size       int
// 	Top        int //栈顶
// }

// //入
// func (s *Stack) Push(str byte) *Stack {
// 	if s.Size > 0 {
// 		s.Top++
// 	}
// 	if s.Size < len(s.StackSlice) { //复用
// 		s.StackSlice[s.Top] = str
// 	} else {
// 		s.StackSlice = append(s.StackSlice, str)
// 	}
// 	s.Size++
// 	return s
// }

// //出
// func (s *Stack) Pop() byte {
// 	var rtnVal byte
// 	if s.Size > 0 {
// 		//fmt.Printf("%v\t", s.StackSlice[s.Top])
// 		rtnVal = s.StackSlice[s.Top]
// 		s.StackSlice[s.Top] = '0'
// 		if s.Top != 0 {
// 			s.Top--
// 		}
// 		s.Size--
// 		return rtnVal
// 	}
// 	return '0'
// }

// //获取栈顶元素但不出栈
// func (s *Stack) GetTopButNoPop() byte {
// 	if s.Size > 0 {
// 		return s.StackSlice[s.Top]
// 	}
// 	return '0'
// }

// @lc code=end
