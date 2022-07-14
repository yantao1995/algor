package leetcode

/*
 * @lc app=leetcode.cn id=745 lang=golang
 *
 * [745] 前缀和后缀搜索
 */

// @lc code=start
type WordFilter struct {
	pm, sm map[string][]int
}

// func Constructor(words []string) WordFilter {
// 	w := WordFilter{
// 		pm: map[string][]int{},
// 		sm: map[string][]int{},
// 	}
// 	for k := range words {
// 		for i := 0; i <= len(words[k]); i++ {
// 			w.pm[words[k][:i]] = append(w.pm[words[k][:i]], k)
// 			w.sm[words[k][i:]] = append(w.sm[words[k][i:]], k)
// 		}
// 	}
// 	return w
// }

func (this *WordFilter) F(pref string, suff string) int {
	ps, ok := this.pm[pref]
	if !ok {
		return -1
	}
	ss, ok := this.sm[suff]
	if !ok {
		return -1
	}
	for i := len(ps) - 1; i >= 0; i-- {
		for j := len(ss) - 1; j >= 0; j-- {
			if ps[i] == ss[j] {
				return ps[i]
			}
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
// @lc code=end

/*
	两个map，pm存前缀，sm存后缀
	Constructor 将每个单词拆解成前缀和后缀放进去
	F 如果找到前缀和后缀都存在，因为题目要求返回最大的，
		就从数组从后向前遍历，找到最大的返回。
*/
