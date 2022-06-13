package leetcode

// func replaceSpace(s string) string {
// 	count := 0
// 	for k := range s {
// 		if s[k] == ' ' {
// 			count++
// 		}
// 	}
// 	bts := make([]byte, 0, len(s)+count*2)
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == ' ' {
// 			bts = append(bts, '%', '2', '0')
// 		} else {
// 			bts = append(bts, s[i])
// 		}
// 	}
// 	return string(bts)
// }

// /*
// 	遍历S，查看需要构建的bts长度是多少，即有几个空格
// 	然后遍历，逐一追加至bts
// */

//双指针
func replaceSpace(s string) string {
	count := 0
	for k := range s {
		if s[k] == ' ' {
			count++
		}
	}
	bts := append([]byte(s), make([]byte, count*2)...)
	for s, e := len(s)-1, len(bts)-1; s >= 0 && e >= 0; {
		if bts[s] == ' ' {
			bts[e] = '0'
			e--
			bts[e] = '2'
			e--
			bts[e] = '%'
		} else {
			bts[e] = bts[s]
		}
		e--
		s--
	}
	return string(bts)
}

/*
	s和e 分别指向了 bts的数据末尾 和 字符串末尾
	从后往前，避免数据后移退化成O(n^2)
	遇到空格就追加3个字符
*/
