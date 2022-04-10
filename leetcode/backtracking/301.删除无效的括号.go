package leetcode

/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	left, right := 0, 0
	for k := range s {
		if s[k] == '(' {
			left++
		} else if s[k] == ')' {
			if left == 0 {
				right++
			} else {
				left--
			}
		}
	}
	check := func(str string) bool {
		count := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '(' {
				count++
			} else if str[i] == ')' {
				count--
				if count < 0 {
					return false
				}
			}
		}
		return count == 0
	}
	distinct := map[string]struct{}{}
	var backtrace func(cl, cr int, str string)
	backtrace = func(cl, cr int, str string) {
		if cl == 0 && cr == 0 {
			if check(str) {
				distinct[string(str)] = struct{}{}
			}
			return
		}
		for i := 0; i < len(str); i++ {
			if cl+cr > len(str)-i { //数量都不够了，后面的操作肯定也都不满足
				return
			}
			if str[i] == '(' && cl > 0 && (i == 0 || str[i] != str[i-1]) { //去掉 类似于 ((  连续两个括号这种，因为删去哪个都一样的效果
				backtrace(cl-1, cr, str[:i]+str[i+1:])
			}
			if str[i] == ')' && cr > 0 && (i == 0 || str[i] != str[i-1]) { //去掉 类似于 )) 连续两个括号这种，因为删去哪个都一样的效果
				backtrace(cl, cr-1, str[:i]+str[i+1:])
			}
		}
	}
	backtrace(left, right, s)
	result := make([]string, 0, len(distinct))
	for k := range distinct {
		result = append(result, k)
	}
	return result
}

// @lc code=end

/*
	回溯法  不剪枝的话会超时
	left,right  记录匹配需要删去的括号数量
	distinct 用于结果集去重
	check() 校验当前字符串是否是符合条件的

	本质就是暴力解法，挨个删了去试，
	当left和right都删完了就去判断是否符合条件，符合就加入到去重map里面
	然后 三个地方进行剪枝
		1.  cl + cr > len(str)-i   		  //数量都不够了，后面的操作肯定也都不满足
		2.  i == 0 || str[i] != str[i-1]  //去掉 类似于 ((  连续两个括号这种，因为删去哪个都一样的效果
		3.  i == 0 || str[i] != str[i-1]  //去掉 类似于 )) 连续两个括号这种，因为删去哪个都一样的效果
*/
