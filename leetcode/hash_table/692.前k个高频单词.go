package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for k := range words {
		m[words[k]]++
	}
	weight := []struct {
		word   string
		weight int
	}{}
	for k, v := range m {
		weight = append(weight, struct {
			word   string
			weight int
		}{k, v})
	}
	sort.Slice(weight, func(i, j int) bool {
		if weight[i].weight != weight[j].weight {
			return weight[i].weight > weight[j].weight
		}
		return weight[i].word < weight[j].word
	})
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = weight[i].word
	}
	return result
}

// @lc code=end

/*
	map存出现的次数，然后对key排序
*/
