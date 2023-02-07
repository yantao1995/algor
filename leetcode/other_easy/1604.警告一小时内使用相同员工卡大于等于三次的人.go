package leetcode

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1604 lang=golang
 *
 * [1604] 警告一小时内使用相同员工卡大于等于三次的人
 */

// @lc code=start
func alertNames(keyName []string, keyTime []string) []string {
	result := []string{}
	isWarn := func(t1, t2 string) bool {
		tm1, _ := strconv.Atoi(t1[:2])
		tm2, _ := strconv.Atoi(t2[:2])
		ts1, _ := strconv.Atoi(t1[3:])
		ts2, _ := strconv.Atoi(t2[3:])
		if tm2 >= tm1 {
			return ts2-ts1+(tm2-tm1)*60 <= 60
		}
		return false
	}
	m := map[string][]string{}
	for k := range keyName {
		m[keyName[k]] = append(m[keyName[k]], keyTime[k])
	}
	for name, times := range m {
		sort.Strings(times)
		for i := 2; i < len(times); i++ {
			if isWarn(times[i-2], times[i]) {
				result = append(result, name)
				break
			}
		}
	}
	sort.Strings(result)
	return result
}

// @lc code=end

/*
	题目与用例描述的不准确，用例全给的是同一天的，排序后计算即可
*/
