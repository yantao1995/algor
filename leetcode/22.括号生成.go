package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
// func generateParenthesis(n int) []string {
// 	all, fit := map[string]bool{}, []string{}
// 	nums := (n - 1) * 2
// 	bts := make([]byte, n*2)
// 	for i := 1; i < nums; i++ {
// 		bts[i] = '('
// 	}
// 	bts[0], bts[nums+1] = '(', ')'
// 	all[string(bts)] = true
// 	//"(()())"
// 	for ll := 1; ll <= nums; ll++ {
// 		for k := range all {
// 			for bl := 0; bl < 2; bl++ {
// 				temp := []byte(k)
// 				if bl == 0 {
// 					temp[ll] = '('
// 				} else {
// 					temp[ll] = ')'
// 				}
// 				all[string(temp)] = true
// 			}
// 		}
// 	}
// 	for s := range all {
// 		stack := Stack{
// 			StackSlice: []byte{},
// 			Size:       0,
// 			Top:        0,
// 		}
// 		for k := range s {
// 			if stack.Size > 0 {
// 				if s[k]-stack.GetTopButNoPop() == 1 {
// 					stack.Pop()
// 				} else {
// 					stack.Push(s[k])
// 				}
// 			} else {
// 				stack.Push(s[k])
// 			}
// 		}
// 		if stack.Size > 0 {
// 			continue
// 		}
// 		fit = append(fit, s)
// 	}
// 	return fit
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
