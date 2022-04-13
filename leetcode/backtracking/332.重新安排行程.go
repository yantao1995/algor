package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	type ft struct {
		from, to string
	}
	trail := map[ft]int{}
	for k := range tickets {
		trail[ft{tickets[k][0], tickets[k][1]}]++
	}
	tkm := map[string][]string{} //航线
	for k := range tickets {
		if _, ok := tkm[tickets[k][0]]; !ok {
			tkm[tickets[k][0]] = []string{}
		}
		for _, v := range tkm[tickets[k][0]] {
			if v == tickets[k][1] {
				goto lab
			}
		}
		tkm[tickets[k][0]] = append(tkm[tickets[k][0]], tickets[k][1])
	lab:
	}
	for k := range tkm {
		sort.Strings(tkm[k])
	}
	result := []string{}
	var backtrace func(from string, count int, ctrail map[ft]int, temp []string) //出发地，用票次数，已跑的轨迹，当前轨迹
	backtrace = func(from string, count int, ctrail map[ft]int, temp []string) {
		if len(result) == len(tickets)+1 || count == len(tickets) {
			if count == len(tickets) {
				result = append(result, temp...)
			}
			return
		}
		for _, to := range tkm[from] {
			if ctrail[ft{from, to}] < trail[ft{from, to}] {
				ctrail[ft{from, to}]++
				backtrace(to, count+1, ctrail, append(temp, to))
				ctrail[ft{from, to}]--
			}
		}

	}
	temp := make([]string, 0, len(tickets)+1)
	backtrace("JFK", 0, map[ft]int{}, append(temp, "JFK"))
	return result
}

// @lc code=end

/*
	tkm 记录航线， 例 A 可以到 B C D    [lab 用于对目标地的去重]
	trail 记录每个 出发地到目的地的次数
	from 表示当前的起点
	count 记录当前已经访问目的地的数量
	ctrail 表示当前已经飞过的航线，与 trail 做判断，在当前票没有用超过次数的情况下进入下轮回溯
	temp 记录当前飞行的轨迹

	从每个起点来依次回溯每个当前可以去的目的的

	剪枝: 如果已经有结果了，即: len(result) == len(tickets)+1 ，那么当前就是最优，可以直接退出了

*/
