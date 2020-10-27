package easy

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	ht1, ht2 := map[string]int{}, map[string]int{}
	tempSlice := []string{}
	for k := 0; k < len(cpdomains); k++ {
		sp1 := strings.Split(cpdomains[k], " ")
		num, _ := strconv.Atoi(sp1[0])
		if ht1[sp1[1]] == 0 {
			tempSlice = append(tempSlice, sp1[1])
			cpdomains[k] = sp1[1]
		}
		ht1[sp1[1]] += num
	}
	cpdomains = tempSlice
	for k, v := range cpdomains {
		ht2[v] += ht1[cpdomains[k]]
		for strings.Contains(v, ".") {
			t2 := strings.SplitN(v, ".", 2)
			v = t2[1]
			ht2[v] += ht1[cpdomains[k]]
		}
	}
	result := []string{}
	for k := range ht2 {
		result = append(result, fmt.Sprintf("%d %s", ht2[k], k))
	}
	return result
}

// @lc code=end
