package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=636 lang=golang
 *
 * [636] 函数的独占时间
 */

// @lc code=start
func exclusiveTime(n int, logs []string) []int {
	type excl struct {
		tm, inner int
	}
	stack := []excl{}
	result := make([]int, n)
	pid, tm, isStart := 0, 0, true
	getSlice := func(k int) {
		slice := strings.Split(logs[k], ":")
		pid, _ = strconv.Atoi(slice[0])
		tm, _ = strconv.Atoi(slice[2])
		isStart = "start" == slice[1]
	}
	topExcl := excl{}
	inner := 0
	for k := range logs {
		getSlice(k)
		if isStart {
			stack = append(stack, excl{
				tm:    tm,
				inner: inner,
			})
		} else {
			topExcl = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp := tm + 1 - topExcl.tm + topExcl.inner - inner
			result[pid] += temp
			inner += temp
		}

	}
	return result
}

// @lc code=end

/*
	栈模拟函数调用
	getSlice 给 单个字符串做拆分
	当前如果是end，那么栈顶必为start与之对应
	topExcl 记录的是当前栈的 内部 走过的任务 总时间 以及当时调用任务的时间
	每次遇到 end ，只需要计算当前 end与 start 的时间差，减去 中间其他任务消耗的时间
*/
